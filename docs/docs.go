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
        "/api/dishes": {
            "get": {
                "description": "Retrieve a list of all available dishes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "Get all dishes",
                "responses": {
                    "200": {
                        "description": "List of dishes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Dish"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "No dishes found",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve dishes",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/dishes/admin/add": {
            "post": {
                "description": "Create a new dish with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes/Secure"
                ],
                "summary": "Add a new dish",
                "parameters": [
                    {
                        "description": "Dish details",
                        "name": "dish",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AddDishPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Dish created successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to add dish",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/dishes/admin/delete/{id}": {
            "delete": {
                "description": "Remove a dish from the system by its ID",
                "tags": [
                    "Dishes/Secure"
                ],
                "summary": "Delete a dish by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dish ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dish deleted successfully"
                    },
                    "400": {
                        "description": "Invalid dish ID",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to delete dish",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/dishes/admin/update": {
            "put": {
                "description": "Update the details of an existing dish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes/Secure"
                ],
                "summary": "Update dish details",
                "parameters": [
                    {
                        "description": "Dish details",
                        "name": "dish",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ChangeDishPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dish updated successfully"
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to update dish",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/dishes/by_category": {
            "post": {
                "description": "Retrieve a list of dishes based on their category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dishes"
                ],
                "summary": "Get dishes by category",
                "parameters": [
                    {
                        "description": "Category details",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.GetDishesByCategoryPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of dishes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Dish"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "404": {
                        "description": "No dishes found",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve dishes",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/dishes/by_id/{dish_id}": {
            "get": {
                "description": "Retrieve a specific dish by its ID",
                "tags": [
                    "Dishes"
                ],
                "summary": "Get a dish by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dish ID",
                        "name": "dish_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dish details",
                        "schema": {
                            "$ref": "#/definitions/models.Dish"
                        }
                    },
                    "400": {
                        "description": "Invalid dish ID",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "404": {
                        "description": "Dish not found",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve dish",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/api/dishes/search/{name}": {
            "get": {
                "description": "Search for dishes by their name",
                "tags": [
                    "Dishes"
                ],
                "summary": "Search dishes by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dish name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of matching dishes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Dish"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Query parameter 'name' is required",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "404": {
                        "description": "No dishes found",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Failed to search dishes",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
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
                        "description": "Updated User Creds",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangeUserCredsPayload"
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
        "handlers.AddDishPayload": {
            "type": "object",
            "required": [
                "dish_category",
                "dish_description",
                "dish_name",
                "dish_photo",
                "dish_price",
                "dish_weight"
            ],
            "properties": {
                "dish_category": {
                    "type": "integer"
                },
                "dish_description": {
                    "type": "string"
                },
                "dish_name": {
                    "type": "string"
                },
                "dish_photo": {
                    "type": "string"
                },
                "dish_price": {
                    "type": "number"
                },
                "dish_weight": {
                    "type": "number"
                }
            }
        },
        "handlers.ChangeDishPayload": {
            "type": "object",
            "required": [
                "dish_category",
                "dish_description",
                "dish_name",
                "dish_photo",
                "dish_price",
                "dish_weight",
                "id"
            ],
            "properties": {
                "dish_category": {
                    "type": "integer"
                },
                "dish_description": {
                    "type": "string"
                },
                "dish_name": {
                    "type": "string"
                },
                "dish_photo": {
                    "type": "string"
                },
                "dish_price": {
                    "type": "number"
                },
                "dish_weight": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
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
        "handlers.GetDishesByCategoryPayload": {
            "type": "object",
            "required": [
                "dish_category"
            ],
            "properties": {
                "dish_category": {
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
        "models.ChangeUserCredsPayload": {
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
                "user_surname": {
                    "type": "string"
                }
            }
        },
        "models.Dish": {
            "type": "object",
            "required": [
                "dish_description",
                "dish_name",
                "dish_photo_url",
                "dish_price",
                "dish_weight"
            ],
            "properties": {
                "dish_category": {
                    "type": "string"
                },
                "dish_description": {
                    "type": "string"
                },
                "dish_name": {
                    "type": "string"
                },
                "dish_photo_url": {
                    "type": "string"
                },
                "dish_price": {
                    "type": "number"
                },
                "dish_rating": {
                    "type": "integer"
                },
                "dish_weight": {
                    "type": "number"
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
                "user_photo",
                "user_role",
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
