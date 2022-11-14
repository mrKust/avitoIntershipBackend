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
        "contact": {
            "email": "dimvas2010@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/billing": {
            "post": {
                "description": "Add money to user's balance with billing systems (visa/mastercard)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "billings"
                ],
                "summary": "Add money to user's balance",
                "parameters": [
                    {
                        "description": "Users Info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/moneyAccept": {
            "post": {
                "description": "Accept money from master balance when service is done",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "reserve"
                ],
                "summary": "Accepts money",
                "parameters": [
                    {
                        "description": "Request to accept money",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/masterBalance.MasterBalance"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/moneyFreeze": {
            "post": {
                "description": "Reserves money from user balance to special master account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "reserve"
                ],
                "summary": "Reserves money",
                "parameters": [
                    {
                        "description": "Request to freeze money",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/masterBalance.MasterBalance"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/moneyReject": {
            "post": {
                "description": "Return money to user when payment for service is rejected",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "reserve",
                    "reject"
                ],
                "summary": "Rejects money",
                "parameters": [
                    {
                        "description": "Request to freeze money",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/masterBalance.MasterBalance"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/report/:month/:year": {
            "get": {
                "description": "Return link to report.csv file with money for every service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "balance",
                    "report"
                ],
                "summary": "Returns report for date range",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Needed month for report",
                        "name": "month",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Needed year for report",
                        "name": "year",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/transactions/:userid/:pageNum/:sortSum/:sortDate": {
            "get": {
                "description": "Return text with history of transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "balance"
                ],
                "summary": "Returns info about user transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of needed user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "number of searching page",
                        "name": "pageNum",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Parameter for sort by sum (asc/desc)",
                        "name": "sortSum",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Parameter for sort by date (asc/desc)",
                        "name": "sortDate",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/:id": {
            "get": {
                "description": "Return user account with his balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts",
                    "balance"
                ],
                "summary": "Returns user balance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "masterBalance.MasterBalance": {
            "type": "object",
            "properties": {
                "from_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "money_amount": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "service_id": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Avito backend internship",
	Description:      "This service help to work with users balance",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
