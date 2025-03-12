package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
	secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret: secret}
}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Отладочная информация
		fmt.Printf("Request path: %s\n", r.URL.Path)

		// Проверка на публичные маршруты
		if r.URL.Path == "/user/login" ||
			r.URL.Path == "/user/register" ||
			strings.HasPrefix(r.URL.Path, "/swagger/") {

			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		fmt.Printf("Auth header: %s\n", authHeader)

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Empty authorization header"))
			return
		}

		authString := strings.Split(authHeader, " ")
		if len(authString) != 2 || authString[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unable to read Authorization header"))
			return
		}

		// Отладочная информация о токене
		tokenStr := authString[1]
		fmt.Printf("Token: %s\n", tokenStr)

		// Парсинг токена
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.secret), nil
		})

		if err != nil {
			fmt.Printf("Token parse error: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("something wrong with token: " + err.Error()))
			return
		}

		// Проверка валидности токена
		if !token.Valid {
			fmt.Println("Token is invalid")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}

		fmt.Println("Token is valid")

		// Получение claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Failed to get claims from token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token claims format"))
			return
		}

		// Отладочная информация о claims
		fmt.Println("Token claims:")
		for k, v := range claims {
			fmt.Printf("  %s: (%T) %v\n", k, v, v)
		}

		// Проверка времени истечения
		expValue, ok := claims["exp"].(float64)
		if !ok {
			fmt.Println("exp claim not found or not a number")
			http.Error(w, "Payload does not have exp value", http.StatusUnauthorized)
			return
		}

		expTime := time.Unix(int64(expValue), 0)
		if time.Now().After(expTime) {
			fmt.Println("Token expired")
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		// Получение user_id
		userIDRaw, ok := claims["user_id"]
		if !ok {
			fmt.Println("user_id claim not found")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing user_id in payload"))
			return
		}

		fmt.Printf("user_id from token: (%T) %v\n", userIDRaw, userIDRaw)

		// Обработка user_id в зависимости от типа
		var userID uuid.UUID
		var userIDErr error

		switch v := userIDRaw.(type) {
		case string:
			userID, userIDErr = uuid.Parse(v)
			fmt.Printf("Parsing string user_id: %s\n", v)
		case map[string]interface{}:
			fmt.Printf("user_id is a map: %v\n", v)
			if str, ok := v["String"].(string); ok {
				userID, userIDErr = uuid.Parse(str)
				fmt.Printf("Parsing string from map: %s\n", str)
			} else {
				userIDErr = fmt.Errorf("could not find String field in map")
			}
		default:
			// Попытка преобразовать в строку
			fmt.Printf("Converting unknown type to string: %v\n", v)
			str := fmt.Sprintf("%v", v)
			userID, userIDErr = uuid.Parse(str)
		}

		if userIDErr != nil {
			fmt.Printf("Error parsing user_id: %v\n", userIDErr)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("Invalid user_id format: %v", userIDErr)))
			return
		}

		fmt.Printf("Successfully parsed user_id as UUID: %s\n", userID.String())

		// Сохраняем в контекст с ключом "userId" для согласованности с обработчиками
		ctx := context.WithValue(r.Context(), "userId", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
