package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w, status: http.StatusBadRequest}
}

func (rw *ResponseWriter) Status() int {
	return rw.status
}

func (rw *ResponseWriter) Size() int {
	return rw.size
}

func (rw *ResponseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		customWriter := NewResponseWriter(w)

		log := fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(customWriter, r)

		duration := time.Since(start)

		slog.Info("Request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", customWriter.Status(),
			"duration", duration,
			"details", log)
	})
}
