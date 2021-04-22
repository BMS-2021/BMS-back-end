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
                "produces": [
                    "text/plain"
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
                        "type": "boolean",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order",
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Book"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "text/csv"
                ],
                "produces": [
                    "text/plain"
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
        "/borrow": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Get the books borrowed by a specific Bard",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Card ID",
                        "name": "cardId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Book"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "text/plain",
                    "application/json"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Borrow a new book",
                "parameters": [
                    {
                        "description": " ",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BorrowReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "The requested book has been all borrowed, return the last borrowed time object",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Book not found or Card not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/card": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Card"
                ],
                "summary": "Retrieve a library card",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Card ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Card"
                ],
                "summary": "Create a library card",
                "parameters": [
                    {
                        "description": " ",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CardReq"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Card"
                ],
                "summary": "Delete a library card",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Card ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/login": {
            "get": {
                "tags": [
                    "Login"
                ],
                "summary": "Check login status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AdminResp"
                        }
                    },
                    "401": {
                        "description": "Not logged in",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Admin login",
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
        },
        "/return": {
            "post": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Borrow"
                ],
                "summary": "Return a book",
                "parameters": [
                    {
                        "description": " ",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BorrowReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": "Book not found or Card not found or Borrow not found",
                        "schema": {
                            "type": "string"
                        }
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
        "model.AdminResp": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
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
        },
        "model.BookReq": {
            "type": "object",
            "required": [
                "author",
                "category",
                "press",
                "price",
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
        },
        "model.BorrowReq": {
            "type": "object",
            "required": [
                "bookId",
                "cardId"
            ],
            "properties": {
                "bookId": {
                    "type": "integer"
                },
                "cardId": {
                    "type": "integer"
                }
            }
        },
        "model.CardReq": {
            "type": "object",
            "required": [
                "department",
                "name",
                "type"
            ],
            "properties": {
                "department": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
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
	Version:     "0.3",
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
