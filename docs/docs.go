// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "info.donilz@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/films": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all films added by users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Film"
                            }
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/films/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves film based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Film ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Film"
                        }
                    },
                    "400": {
                        "description": "Invalid film ID",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "404": {
                        "description": "Film with specified ID not found",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/funcs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all functions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/funcs/{funcName}": {
            "put": {
                "summary": "Call function based on given funcName",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Function name",
                        "name": "funcName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Function successfully called",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "404": {
                        "description": "Function with specified funcName not found",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all registered users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserInfo"
                            }
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a new user",
                "responses": {
                    "200": {
                        "description": "Registration completed successfully",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid register data",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "409": {
                        "description": "User with such data is already registered",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/users/{userName}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves user based on given UserName (Login)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserName (Login)",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/{userName}/films": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all films added by specified user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserName (Login)",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Film"
                            }
                        }
                    },
                    "401": {
                        "description": "Not authorized",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adds a new film to the list of the specified user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserName (Login)",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Film successfully added",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "400": {
                        "description": "Incorrect json ((insufficient or incorrect data) or invalid format)",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "401": {
                        "description": "Not authorized",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "403": {
                        "description": "The username in the parameters does not match the name of the authorized user",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "409": {
                        "description": "Film already added",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Removes the specified film from the user's film list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserName (Login)",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Film successfully deleted",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "400": {
                        "description": "Incorrect json ((insufficient or incorrect data) or invalid format)",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "401": {
                        "description": "Not authorized",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "403": {
                        "description": "The username in the parameters does not match the name of the authorized user",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "404": {
                        "description": "Removable film not found",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/model.DefaultResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DefaultResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "model.Film": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Moviefan Swagger API",
	Description: "Swagger API for Golang Project Moviefan",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
