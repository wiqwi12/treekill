// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "support@studynoteapi.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/notes": {
            "get": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Get list of all user's notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Get all notes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Note"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Create new note for authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Create new note",
                "parameters": [
                    {
                        "description": "Note data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNoteRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/notes/{id}": {
            "get": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Get single note by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Get note by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Note ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Update existing note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Update note",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Note ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated note data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateNoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Delete note by ID",
                "tags": [
                    "Notes"
                ],
                "summary": "Delete note",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Note ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Get JWT access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User authentication",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Create new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "description": "Registration data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AuthResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                }
            }
        },
        "dto.CreateNoteRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "Note content here"
                },
                "title": {
                    "type": "string",
                    "example": "My First Note"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd!"
                }
            }
        },
        "dto.RegistrationRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "P@ssw0rd!"
                },
                "username": {
                    "type": "string",
                    "example": "john_doe"
                }
            }
        },
        "dto.UpdateNoteRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "Updated note content"
                },
                "title": {
                    "type": "string",
                    "example": "Updated Note Title"
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "userid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "username": {
                    "type": "string",
                    "example": "john_doe"
                }
            }
        },
        "errors.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error description"
                }
            }
        },
        "models.Note": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "titel": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWTAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.8.2",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "StudyNoteAPI",
	Description:      "API for notes management with JWT authentication",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
