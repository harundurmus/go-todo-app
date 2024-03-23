// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/todo": {
            "post": {
                "description": "Create Or Update todo data",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Or Update todo data",
                "parameters": [
                    {
                        "description": "Todo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/todo.Todo"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "todo.Todo": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "task": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}