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
            "url": "https://github.com/YuStep/",
            "email": "support@mail.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/YuStep/wbl0/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/order": {
            "get": {
                "description": "get all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get All Orders",
                "operationId": "get-all-orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.getAllOrdersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/order/:id": {
            "get": {
                "description": "get order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get Order By Id",
                "operationId": "get-order-by-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.getAllOrdersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                }
            }
        },
        "models.Delivery": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "chrt_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nm_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "rid": {
                    "type": "string"
                },
                "sale": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "number"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "delivery": {
                    "$ref": "#/definitions/models.Delivery"
                },
                "delivery_service": {
                    "type": "string"
                },
                "entry": {
                    "type": "string"
                },
                "internal_signature": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "locale": {
                    "type": "string"
                },
                "oof_shard": {
                    "type": "string"
                },
                "order_uid": {
                    "type": "string"
                },
                "payment": {
                    "$ref": "#/definitions/models.Payment"
                },
                "shard_key": {
                    "type": "string"
                },
                "sm_id": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "bank": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "custom_fee": {
                    "type": "integer"
                },
                "delivery_cost": {
                    "type": "integer"
                },
                "goods_total": {
                    "type": "number"
                },
                "payment_dt": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go WB API",
	Description:      "This is a simple RESTful Service API written in Go using Gin web framework, PostgreSQL database,NATS queue with in-memory caching.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
