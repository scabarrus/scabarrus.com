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
            "name": "scabarrus",
            "email": "scabarrus@gmail.com"
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
        "/groups": {
            "get": {
                "description": "get all groups",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Show all groups",
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Create a group",
                "parameters": [
                    {
                        "description": "dto",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/groups/{group}": {
            "get": {
                "description": "get a group by it's name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Show a group details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "modify a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Modify a group",
                "parameters": [
                    {
                        "description": "dto",
                        "name": "groupdto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Delete a group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GroupDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/groups/{group}/members": {
            "get": {
                "description": "find all roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Find all members",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GroupDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a member's group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Create a member's group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "dto",
                        "name": "user",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/service.GroupMemberService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/groups/{group}/members/{member}": {
            "get": {
                "description": "find a member's group by it's name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Find a member's group by it's name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "member name",
                        "name": "member",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a member's group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Delete a member's group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group name",
                        "name": "group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "member name",
                        "name": "member",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "send OK if it database connexion works",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthz"
                ],
                "summary": "healthcheck",
                "responses": {
                    "200": {
                        "description": "status",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/roles": {
            "get": {
                "description": "get all roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Show all roles",
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Create a role",
                "parameters": [
                    {
                        "description": "dto",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/roles/{role}": {
            "get": {
                "description": "get a role by it's name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Show a role details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "modify a role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Modify a role details",
                "parameters": [
                    {
                        "description": "dto",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Delete a role details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/roles/{role}/members": {
            "get": {
                "description": "find all roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Find all roles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RoleDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a member's group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Create a member's group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "member",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RoleMemberService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "create a member's group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Create a member's group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/roles/{role}/members/{member}": {
            "get": {
                "description": "find a member's role by it's name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Find a member's role by it's name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role name",
                        "name": "role",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "member name",
                        "name": "member",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.RoleDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Show all users",
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "dto",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/users/{user}": {
            "get": {
                "description": "get a user by it's name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Show a user details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "modify a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Modify a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "dto",
                        "name": "userDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user name",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "dto",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.GroupDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Operator group"
                },
                "gid": {
                    "type": "integer",
                    "example": 7001
                },
                "group": {
                    "type": "string",
                    "example": "operator"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.UserDTO"
                    }
                }
            }
        },
        "dto.RoleDTO": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GroupDTO"
                    }
                },
                "namespace": {
                    "type": "string"
                },
                "resource": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "verb": {
                    "type": "string"
                }
            }
        },
        "dto.UserDTO": {
            "type": "object",
            "properties": {
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GroupDTO"
                    }
                },
                "password": {
                    "type": "string",
                    "example": "B67zuopX#2"
                },
                "uid": {
                    "type": "integer",
                    "example": 5000
                },
                "user": {
                    "type": "string",
                    "example": "user1"
                }
            }
        },
        "error.Error": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "service.GroupMemberService": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "operator"
                },
                "user": {
                    "type": "string",
                    "example": "user1"
                }
            }
        },
        "service.RoleMemberService": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "dev"
                },
                "role": {
                    "type": "string",
                    "example": "role1"
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
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "User-Management API",
	Description: "This is a sample serice for managing user and role for kubernetes cluster",
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
