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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/book": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create a single book item",
                "parameters": [
                    {
                        "description": "Information of a book",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BookReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/books": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Retrieve multiple books under multiple conditions",
                "parameters": [
                    {
                        "type": "string",
                        "name": "author",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "press",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "priceMax",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "priceMin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "yearMax",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "yearMin",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "put": {
                "consumes": [
                    "text/csv"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create multiple book items by uploading a csv file",
                "parameters": [
                    {
                        "description": "A csv file, with book information inside",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "parameters": [
                    {
                        "description": "Login information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AdminReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AdminReq": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.BookReq": {
            "type": "object",
            "required": [
                "author",
                "category",
                "press",
                "price",
                "stock",
                "title",
                "total",
                "year"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "press": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
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
	Version:     "0.1",
	Host:        "ralxyz.dev.zjuqsc.com",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Book Management System API",
	Description: "This API will be used under staging environment.",
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
