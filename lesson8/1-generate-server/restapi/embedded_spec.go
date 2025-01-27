// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Демонстрационное API сервиса 'список задач'",
    "title": "Список задач",
    "version": "0.1.2"
  },
  "host": "localhost",
  "basePath": "/api",
  "paths": {
    "/task": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Получение списка задач",
        "operationId": "getTasks",
        "parameters": [
          {
            "type": "boolean",
            "description": "Признак получения уже выполенных задач",
            "name": "show_done",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "400": {
            "description": "Некорректный ввод"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Добавить новую задачу",
        "operationId": "createTask",
        "parameters": [
          {
            "name": "task",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Некорректный ввод"
          },
          "422": {
            "description": "Недопустимые значения"
          }
        }
      }
    },
    "/task/{taskId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Получение задачи по идентификатору",
        "operationId": "getTaskById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Идентификатор задачи",
            "name": "taskId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "400": {
            "description": "Некорректный ввод"
          },
          "404": {
            "description": "Такой задачи нет"
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Изменение задачи",
        "operationId": "updateTask",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Идентификатор задачи",
            "name": "taskId",
            "in": "path",
            "required": true
          },
          {
            "name": "task",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Некорректный ввод"
          },
          "404": {
            "description": "Такой задачи нет"
          },
          "422": {
            "description": "Недопустимые значения"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "summary": "Удаление задачи",
        "operationId": "deletePet",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Идентификатор задачи",
            "name": "taskId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Некорректный ввод"
          },
          "404": {
            "description": "Такой задачи нет"
          }
        }
      }
    }
  },
  "definitions": {
    "Task": {
      "type": "object",
      "required": [
        "id",
        "title"
      ],
      "properties": {
        "description": {
          "description": "Подробное описание",
          "type": "string",
          "example": "Нужен цельнозерновой для диеты"
        },
        "id": {
          "description": "Идентификатор",
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "status": {
          "description": "Статус",
          "type": "string",
          "enum": [
            "active",
            "done"
          ]
        },
        "title": {
          "description": "Название",
          "type": "string",
          "example": "Купить хлеба"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Демонстрационное API сервиса 'список задач'",
    "title": "Список задач",
    "version": "0.1.2"
  },
  "host": "localhost",
  "basePath": "/api",
  "paths": {
    "/task": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Получение списка задач",
        "operationId": "getTasks",
        "parameters": [
          {
            "type": "boolean",
            "description": "Признак получения уже выполенных задач",
            "name": "show_done",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "400": {
            "description": "Некорректный ввод"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Добавить новую задачу",
        "operationId": "createTask",
        "parameters": [
          {
            "name": "task",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Некорректный ввод"
          },
          "422": {
            "description": "Недопустимые значения"
          }
        }
      }
    },
    "/task/{taskId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Получение задачи по идентификатору",
        "operationId": "getTaskById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Идентификатор задачи",
            "name": "taskId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "400": {
            "description": "Некорректный ввод"
          },
          "404": {
            "description": "Такой задачи нет"
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Изменение задачи",
        "operationId": "updateTask",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Идентификатор задачи",
            "name": "taskId",
            "in": "path",
            "required": true
          },
          {
            "name": "task",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Некорректный ввод"
          },
          "404": {
            "description": "Такой задачи нет"
          },
          "422": {
            "description": "Недопустимые значения"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "summary": "Удаление задачи",
        "operationId": "deletePet",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Идентификатор задачи",
            "name": "taskId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Некорректный ввод"
          },
          "404": {
            "description": "Такой задачи нет"
          }
        }
      }
    }
  },
  "definitions": {
    "Task": {
      "type": "object",
      "required": [
        "id",
        "title"
      ],
      "properties": {
        "description": {
          "description": "Подробное описание",
          "type": "string",
          "example": "Нужен цельнозерновой для диеты"
        },
        "id": {
          "description": "Идентификатор",
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "status": {
          "description": "Статус",
          "type": "string",
          "enum": [
            "active",
            "done"
          ]
        },
        "title": {
          "description": "Название",
          "type": "string",
          "example": "Купить хлеба"
        }
      }
    }
  }
}`))
}
