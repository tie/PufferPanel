// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-28 15:22:25.7634743 -0500 CDT m=+20.878736301

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
            "name": "PufferPanel",
            "url": "https://pufferpanel.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/nodes": {
            "get": {
                "description": "Gets all nodes registered to the panel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get nodes",
                "responses": {
                    "200": {
                        "description": "Nodes",
                        "schema": {
                            "$ref": "#/definitions/models.NodesView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Node created",
                        "schema": {
                            "$ref": "#/definitions/models.NodeView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/nodes/{id}": {
            "get": {
                "description": "Gets information about a single node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get node",
                "responses": {
                    "200": {
                        "description": "Nods",
                        "schema": {
                            "$ref": "#/definitions/models.NodeView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates a node with given information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.Empty"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes the node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes a node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.Empty"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/templates": {
            "get": {
                "description": "Gets all templates registered",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get templates",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Templates"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apufferi.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "apufferi.Execution": {
            "type": "object",
            "properties": {
                "arguments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "autorecover": {
                    "type": "boolean"
                },
                "autorestart": {
                    "type": "boolean"
                },
                "autostart": {
                    "type": "boolean"
                },
                "disabled": {
                    "type": "boolean"
                },
                "environmentVars": {
                    "type": "object"
                },
                "post": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apufferi.TypeWithMetadata"
                    }
                },
                "pre": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apufferi.TypeWithMetadata"
                    }
                },
                "program": {
                    "type": "string"
                },
                "stop": {
                    "type": "string"
                },
                "stopCode": {
                    "type": "integer"
                }
            }
        },
        "apufferi.TypeWithMetadata": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "models.NodeView": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "privateHost": {
                    "type": "string"
                },
                "privatePort": {
                    "type": "integer"
                },
                "publicHost": {
                    "type": "string"
                },
                "publicPort": {
                    "type": "integer"
                },
                "sftpPort": {
                    "type": "integer"
                }
            }
        },
        "models.NodesView": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "id": {
                        "type": "integer"
                    },
                    "name": {
                        "type": "string"
                    },
                    "privateHost": {
                        "type": "string"
                    },
                    "privatePort": {
                        "type": "integer"
                    },
                    "publicHost": {
                        "type": "string"
                    },
                    "publicPort": {
                        "type": "integer"
                    },
                    "sftpPort": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.Template": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "display": {
                    "type": "string"
                },
                "environment": {
                    "type": "object",
                    "$ref": "#/definitions/apufferi.TypeWithMetadata"
                },
                "id": {
                    "type": "string"
                },
                "install": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apufferi.TypeWithMetadata"
                    }
                },
                "name": {
                    "type": "string"
                },
                "readme": {
                    "type": "string"
                },
                "run": {
                    "type": "object",
                    "$ref": "#/definitions/apufferi.Execution"
                },
                "supportedEnvironments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apufferi.TypeWithMetadata"
                    }
                },
                "type": {
                    "type": "string"
                },
                "uninstall": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apufferi.TypeWithMetadata"
                    }
                }
            }
        },
        "models.Templates": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "data": {
                        "type": "object"
                    },
                    "display": {
                        "type": "string"
                    },
                    "environment": {
                        "type": "object",
                        "$ref": "#/definitions/apufferi.TypeWithMetadata"
                    },
                    "id": {
                        "type": "string"
                    },
                    "install": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/apufferi.TypeWithMetadata"
                        }
                    },
                    "name": {
                        "type": "string"
                    },
                    "readme": {
                        "type": "string"
                    },
                    "run": {
                        "type": "object",
                        "$ref": "#/definitions/apufferi.Execution"
                    },
                    "supportedEnvironments": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/apufferi.TypeWithMetadata"
                        }
                    },
                    "type": {
                        "type": "string"
                    },
                    "uninstall": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/apufferi.TypeWithMetadata"
                        }
                    }
                }
            }
        },
        "response.Empty": {
            "type": "object"
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "$ref": "#/definitions/apufferi.Error"
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
	Version:     "2.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "PufferPanel API",
	Description: "PufferPanel web interface",
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
