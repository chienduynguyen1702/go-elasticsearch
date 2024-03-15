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
        "/lecturer/": {
            "get": {
                "description": "Get all Lecturer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lecturer"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Lecturer"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new lecturer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lecturer"
                ],
                "parameters": [
                    {
                        "description": "Lecturer object that needs to be added",
                        "name": "lecturer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Lecturer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lecturer created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/lecturer/{document_id}": {
            "get": {
                "description": "Get lecturer by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lecturer"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the lecturer to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Lecturer"
                        }
                    }
                }
            },
            "put": {
                "description": "Update lecturer by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lecturer"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the lecturer to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Lecturer object that needs to be updated",
                        "name": "lecturer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Lecturer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lecturer updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete lecturer by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lecturer"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the lecturer to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lecturer deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/student/": {
            "get": {
                "description": "Get all Student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Student"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "parameters": [
                    {
                        "description": "Student object that needs to be added",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Student created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/student/{document_id}": {
            "get": {
                "description": "Get student by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the student to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                }
            },
            "put": {
                "description": "Update student by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the student to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Student object that needs to be updated",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Student updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete student by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the student to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Student deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/subject/": {
            "get": {
                "description": "Get all Subject",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Subject"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new subject",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "parameters": [
                    {
                        "description": "Subject object that needs to be added",
                        "name": "subject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Subject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subject created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/subject/{document_id}": {
            "get": {
                "description": "Get subject by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the subject to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Subject"
                        }
                    }
                }
            },
            "put": {
                "description": "Update subject by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the subject to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Subject object that needs to be updated",
                        "name": "subject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Subject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subject updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete subject by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subject"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "document_id of the subject to be deleted",
                        "name": "document_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subject deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Lecturer": {
            "type": "object",
            "properties": {
                "lecturer_id": {
                    "type": "string"
                },
                "lecturer_name": {
                    "type": "string"
                }
            }
        },
        "model.Student": {
            "type": "object",
            "properties": {
                "student_id": {
                    "type": "string"
                },
                "student_name": {
                    "type": "string"
                },
                "year_started": {
                    "type": "integer"
                }
            }
        },
        "model.Subject": {
            "type": "object",
            "properties": {
                "subject_id": {
                    "type": "string"
                },
                "subject_name": {
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
