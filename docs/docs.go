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
            "email": "support@example.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/user/change": {
            "patch": {
                "description": "Allows the logged-in user to update their details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user credentials",
                "parameters": [
                    {
                        "description": "Updated User Info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/user/change_password": {
            "patch": {
                "description": "Allows the logged-in user to change their password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Change user password",
                "parameters": [
                    {
                        "description": "Old and New Password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ChangePasswordPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password changed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "401": {
                        "description": "Invalid old password",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/user/delete": {
            "delete": {
                "description": "Deletes the logged-in user's account",
                "tags": [
                    "user"
                ],
                "summary": "Delete user account",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/user/info": {
            "get": {
                "description": "Retrieves the details of the logged-in user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "User information",
                        "schema": {
                            "$ref": "#/definitions/models.UserInfo"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "post": {
                "description": "Logs out the currently logged-in user by clearing the authentication token cookie.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user/photo": {
            "patch": {
                "description": "Allows the logged-in user to update their profile photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user photo",
                "parameters": [
                    {
                        "description": "Photo Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ChangePhotoPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Logs in a user and returns a token in a cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Log in a user",
                "parameters": [
                    {
                        "description": "User Login Credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registers a new user and returns a token in a cookie",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Registration",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserReg"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token in cookie"
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ChangePasswordPayload": {
            "type": "object",
            "required": [
                "new_password",
                "old_password"
            ],
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "handlers.ChangePhotoPayload": {
            "type": "object",
            "required": [
                "photo"
            ],
            "properties": {
                "photo": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginPayload": {
            "type": "object",
            "required": [
                "user_login",
                "user_password"
            ],
            "properties": {
                "user_login": {
                    "type": "string"
                },
                "user_password": {
                    "type": "string"
                }
            }
        },
        "models.APIError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UserInfo": {
            "type": "object",
            "required": [
                "user_address",
                "user_login",
                "user_name",
                "user_phone",
                "user_surname"
            ],
            "properties": {
                "user_address": {
                    "type": "string"
                },
                "user_login": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_phone": {
                    "type": "string"
                },
                "user_photo": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string"
                },
                "user_surname": {
                    "type": "string"
                }
            }
        },
        "models.UserReg": {
            "type": "object",
            "required": [
                "user_address",
                "user_login",
                "user_name",
                "user_password",
                "user_phone",
                "user_surname"
            ],
            "properties": {
                "user_address": {
                    "type": "string"
                },
                "user_login": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_password": {
                    "type": "string"
                },
                "user_phone": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string"
                },
                "user_surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1317",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Delivery Backend API",
	Description:      "API documentation for the Delivery Backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
