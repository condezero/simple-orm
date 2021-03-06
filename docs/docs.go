// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hc": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get all products from db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get all Products.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Adds the product to db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Add Products.",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "AddProductRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/products.AddProductRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get a product from db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get a Product.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update Products.",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "UpdateProductRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/products.UpdateProductRequestBody"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a product from db.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Remove Products.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        "models.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "the id of the product",
                    "type": "integer"
                },
                "name": {
                    "description": "the name of the product",
                    "type": "string"
                },
                "price": {
                    "description": "the price of the product",
                    "type": "integer"
                },
                "stock": {
                    "description": "the stock of the product",
                    "type": "integer"
                }
            }
        },
        "products.AddProductRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name of the product",
                    "type": "string"
                },
                "price": {
                    "description": "Price of the product",
                    "type": "integer"
                },
                "stock": {
                    "description": "Stock of the product",
                    "type": "integer"
                }
            }
        },
        "products.UpdateProductRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name of the product",
                    "type": "string"
                },
                "price": {
                    "description": "Price of the product",
                    "type": "integer"
                },
                "stock": {
                    "description": "Stock of the product",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Simple ORM API",
	Description:      "This is a PoC SPI.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
