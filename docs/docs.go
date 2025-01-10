// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "The Terms of Service for the API http://swagger.io/terms/.",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/book/list": {
            "get": {
                "description": "get book list from title, author or categoyies then return as json message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "get book list by filter",
                "parameters": [
                    {
                        "type": "string",
                        "default": "The Witcher Blood of Elves",
                        "description": "string of book title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "Andrzej Sapkowski",
                        "description": "string of author",
                        "name": "author",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "integer of categories id between 1-4",
                        "name": "categories",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services_book.responseBookList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/book/popular": {
            "get": {
                "description": "find most popular book then return as json responseBook model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "find most popular book",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services_book.responseBook"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/borrow/book/{id}": {
            "put": {
                "description": "changing isBorrow status true is mean borrowed this route will increase borrowed_count plus 1 then return json message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "changing isBorrow status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string id of book",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponseModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/create/book/": {
            "post": {
                "description": "create book by json body as models_book.BookRequest then return as json message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "create book by json body",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Some book title",
                        "description": "string of book title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "my self",
                        "description": "string of book author",
                        "name": "author",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "category fill as number between 1-4",
                        "name": "categories",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "constant of borrow status true is mean borrowed",
                        "name": "is_borrow",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "borrowed counter",
                        "name": "borrowed_count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponseModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/delete/book/{id}": {
            "delete": {
                "description": "delete book by id then return as message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "delete book by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string id of book",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponseModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/lotto/list": {
            "get": {
                "description": "get lotto list by LottoFilterReq model then return lotto list as json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lotties"
                ],
                "summary": "Response lotto list by filter",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2025",
                        "description": "year as string number",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "200",
                        "description": "3 digit of string number",
                        "name": "b3d",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "20",
                        "description": "2 digit of string number",
                        "name": "b2d",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "200",
                        "description": "3 digit of string number",
                        "name": "f3d",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "730209",
                        "description": "6 digit of string number",
                        "name": "first",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services_lotto.responseLottoList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/lotto/{id}": {
            "get": {
                "description": "get string by id then return lotto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lotties"
                ],
                "summary": "Show an lotto",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exmaple : 677accbf56ff1bc2a731b944",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services_lotto.responseLotto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/movie/list": {
            "get": {
                "description": "get movie list by filter model then return movie list as json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Response movie list by filter",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services_movie.responseMovieList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/movie/{id}": {
            "get": {
                "description": "get movie by id param then return movie detail as json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Response movie by id",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2025",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services_movie.responseMovie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/update/book/{id}": {
            "put": {
                "description": "update book from param id from _id then return as json message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "update book from param id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Example : 678139d4e150ae9affc077bf",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Some book title",
                        "description": "string of book title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "my self",
                        "description": "string of book author",
                        "name": "author",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "category fill as number between 1-4",
                        "name": "categories",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "constant of borrow status true is mean borrowed",
                        "name": "is_borrow",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "borrowed counter",
                        "name": "borrowed_count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponseModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ErrorResponseModel": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Error Message"
                }
            }
        },
        "handler.SuccessResponseModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Success Message"
                }
            }
        },
        "models_book.BookFilterResponse": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "borrowed_count": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_borrow": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models_lotto.Lotto": {
            "type": "object",
            "properties": {
                "back_three_digit_prize": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['376'",
                        " '297']"
                    ]
                },
                "back_two_digit_prize": {
                    "type": "string",
                    "example": "51"
                },
                "date": {
                    "type": "string",
                    "example": "2025-01-01T00:00:00Z"
                },
                "first_prize": {
                    "type": "string",
                    "example": "730209"
                },
                "front_three_digit_prize": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['446'",
                        " '065']"
                    ]
                }
            }
        },
        "models_movie.MovieRes": {
            "type": "object",
            "properties": {
                "adult": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "services_book.responseBook": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models_book.BookFilterResponse"
                }
            }
        },
        "services_book.responseBookList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models_book.BookFilterResponse"
                    }
                }
            }
        },
        "services_lotto.responseLotto": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models_lotto.Lotto"
                }
            }
        },
        "services_lotto.responseLottoList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models_lotto.Lotto"
                    }
                }
            }
        },
        "services_movie.responseMovie": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models_movie.MovieRes"
                }
            }
        },
        "services_movie.responseMovieList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models_movie.MovieRes"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.1",
	Host:             "tr-profile-go-web-server.onrender.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "REST API using GO and Gin by TR",
	Description:      "This is a API Swagger documentation included book, movie and lotto API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
