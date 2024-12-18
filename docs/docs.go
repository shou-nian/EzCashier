// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin": {
            "put": {
                "description": "Update the role of an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user role",
                "parameters": [
                    {
                        "description": "User role update data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User role updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete an existing user",
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
                        "description": "User deletion data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DeleteUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate a user and return a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logout a user and invalidate their JWT token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "Successful logout",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/password": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update the password of the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user password",
                "parameters": [
                    {
                        "description": "Password update data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated successfully",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update the information of the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user information",
                "parameters": [
                    {
                        "description": "User info update data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User information updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User creation data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.CreateUserRequest": {
            "type": "object",
            "required": [
                "name",
                "password",
                "phone_num",
                "role"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 8
                },
                "phone_num": {
                    "type": "string"
                },
                "role": {
                    "enum": [
                        "admin",
                        "user",
                        "viewer"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.UserRole"
                        }
                    ]
                }
            }
        },
        "models.DeleteUser": {
            "type": "object",
            "required": [
                "phone_num"
            ],
            "properties": {
                "phone_num": {
                    "type": "string"
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "phone_num"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 8
                },
                "phone_num": {
                    "type": "string"
                }
            }
        },
        "models.UpdatePassword": {
            "type": "object",
            "required": [
                "confirm_password",
                "new_password",
                "old_password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 8
                },
                "new_password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 8
                },
                "old_password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 8
                }
            }
        },
        "models.UpdateUserInfoRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                },
                "phone_num": {
                    "type": "string"
                }
            }
        },
        "models.UpdateUserRoleRequest": {
            "type": "object",
            "required": [
                "phone_num",
                "role"
            ],
            "properties": {
                "phone_num": {
                    "type": "string"
                },
                "role": {
                    "enum": [
                        "admin",
                        "user",
                        "viewer"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.UserRole"
                        }
                    ]
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "name",
                "phone_num"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "Name is the full name of the user",
                    "type": "string"
                },
                "phone_num": {
                    "description": "Email is the unique identifier for the user",
                    "type": "string"
                },
                "role": {
                    "description": "Role defines the user's role in the system, determining their permissions",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.UserRole"
                        }
                    ]
                },
                "status": {
                    "description": "UserStatus represents the current status of the user (e.g., active, inactive)",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.UserStatus"
                        }
                    ]
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.UserRole": {
            "type": "string",
            "enum": [
                "admin",
                "user",
                "viewer"
            ],
            "x-enum-varnames": [
                "RoleAdmin",
                "RoleUser",
                "RoleViewer"
            ]
        },
        "models.UserStatus": {
            "type": "string",
            "enum": [
                "active",
                "inactive"
            ],
            "x-enum-varnames": [
                "StatusActive",
                "StatusInactive"
            ]
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "EzCashier API",
	Description:      "This is the API server for the EzCashier application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
