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
  "swagger": "2.0",
  "info": {
    "title": "OpeinAPI specs for test project",
    "version": "1.0"
  },
  "basePath": "/v1",
  "paths": {
    "/pizzas": {
      "get": {
        "tags": [
          "pizzas"
        ],
        "summary": "Get all the pizzas",
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid request"
          },
          "default": {
            "description": "Error"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "operations about pizzas",
      "name": "pizzas"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "OpeinAPI specs for test project",
    "version": "1.0"
  },
  "basePath": "/v1",
  "paths": {
    "/pizzas": {
      "get": {
        "tags": [
          "pizzas"
        ],
        "summary": "Get all the pizzas",
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid request"
          },
          "default": {
            "description": "Error"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "operations about pizzas",
      "name": "pizzas"
    }
  ]
}`))
}
