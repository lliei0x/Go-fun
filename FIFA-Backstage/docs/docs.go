// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-17 22:50:23.6550711 +0800 CST m=+0.064959901

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
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
        "version": "1.0"
    },
    "paths": {
        "/matches/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Match"
                ],
                "summary": "List Match",
                "parameters": [
                    {
                        "type": "string",
                        "description": "matchID",
                        "name": "matchID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.MatchSerializer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.MatchSerializer": {
            "type": "object",
            "properties": {
                "country_left": {
                    "type": "string"
                },
                "country_right": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "delete_at": {
                    "type": "string"
                },
                "group_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "match_number": {
                    "type": "integer"
                },
                "score": {
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
