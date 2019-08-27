// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-27 13:46:33.623228 +0200 CEST m=+0.055910936

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "This is a rest api for basket which saves items to redis server",
        "title": "Basket API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/items": {
            "post": {
                "description": "add by json new CustomerBasket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CustomerBasket"
                ],
                "summary": "Add a CustomerBasket",
                "parameters": [
                    {
                        "description": "Add CustomerBasket",
                        "name": "CustomerBasket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.CustomerBasket"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.CustomerBasket"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "description": "Get CustomerBasket by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CustomerBasket"
                ],
                "summary": "Gets a CustomerBasket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CustomerBasket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.CustomerBasket"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BasketItem": {
            "type": "object",
            "properties": {
                "foodId": {
                    "type": "string"
                },
                "foodName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "oldUnitPrice": {
                    "type": "number"
                },
                "picture": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "unitPrice": {
                    "type": "number"
                }
            }
        },
        "models.CustomerBasket": {
            "type": "object",
            "properties": {
                "customerId": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BasketItem"
                    }
                }
            }
        },
        "models.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    },
    "securityDefinitions": {
        "Identity Server OAuth": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "http://localhost/identity/connect/authorize"
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
