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
        "contact": {
            "name": "Rodrigo Santos",
            "email": "ro.augusto@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/funds": {
            "post": {
                "description": "Process Funds of customers from a specific file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Records"
                ],
                "summary": "Process Funds from a specific file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "input.txt",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/funds-body-req": {
            "post": {
                "description": "Process Funds of customers from body json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Records"
                ],
                "summary": "Process Funds from body json",
                "parameters": [
                    {
                        "description": "List of Load Funds of customers",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/records.RecordAccountList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/funds-write-result-db": {
            "post": {
                "description": "Process Funds of customers from a specific file and write result on MongoDB",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Records"
                ],
                "summary": "Process Funds from a specific file and write result on MongoDB",
                "parameters": [
                    {
                        "type": "file",
                        "description": "input.txt",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/funds/download": {
            "get": {
                "description": "Download the last result of loading the file of Load Funds of customers",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Records"
                ],
                "summary": "Download last result of load funds file",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/funds/result": {
            "get": {
                "description": "List the last result of loading the file of Load Funds of customers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Records"
                ],
                "summary": "List the last result of load funds file",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/records.RecordProcessedList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "records.RecordAccount": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string",
                    "example": "4567"
                },
                "id": {
                    "type": "string",
                    "example": "1234"
                },
                "load_amount": {
                    "type": "string",
                    "example": "$3456.78"
                },
                "time": {
                    "type": "string",
                    "example": "2000-01-01T00:00:00Z"
                }
            }
        },
        "records.RecordAccountList": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/records.RecordAccount"
            }
        },
        "records.RecordProcessed": {
            "type": "object",
            "properties": {
                "accepted": {
                    "type": "string"
                },
                "cod_error": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "load_amount": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "records.RecordProcessedList": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/records.RecordProcessed"
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
	Host:        "localhost:3333",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Koho Load Funds",
	Description: "API REST to load funds based on a file with load amounts.",
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
