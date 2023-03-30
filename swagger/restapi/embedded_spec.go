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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Swagger",
    "title": "Swagger",
    "version": "1.0.0"
  },
  "paths": {
    "/administrator/protected": {
      "$ref": "./prompt/protected_for_administrator.yml"
    },
    "/employee/protected": {
      "$ref": "./prompt/protected_for_employee.yml"
    },
    "/employee/salary_statement": {
      "$ref": "./salary_statement/get_for_employee.yml"
    },
    "/unprotected": {
      "$ref": "./prompt/unprotected.yml"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Swagger",
    "title": "Swagger",
    "version": "1.0.0"
  },
  "paths": {
    "/administrator/protected": {
      "get": {
        "description": "登録済みの管理者のみ実行可能なお試しAPI",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/employee/protected": {
      "get": {
        "description": "登録済みの従業員のみ実行可能なお試しAPI",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/employee/salary_statement": {
      "get": {
        "description": "特定の年月の給料明細を取得するための従業員向けAPI",
        "parameters": [
          {
            "maximum": 3000,
            "minimum": 1500,
            "type": "integer",
            "format": "int",
            "description": "欲しい給料明細の年",
            "name": "year",
            "in": "query",
            "required": true
          },
          {
            "maximum": 12,
            "minimum": 1,
            "type": "integer",
            "format": "int32",
            "description": "欲しい給料明細の月",
            "name": "month",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "amount_of_deduction": {
                  "description": "控除の総額",
                  "type": "integer",
                  "format": "int32"
                },
                "amount_of_earning": {
                  "description": "支給の総額",
                  "type": "integer",
                  "format": "int32"
                },
                "deduction_details": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/DeductionDetailsItems0"
                  }
                },
                "earning_details": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/EarningDetailsItems0"
                  }
                },
                "name_of_employee": {
                  "description": "従業員名",
                  "type": "string"
                },
                "payday": {
                  "description": "給料支払い日時",
                  "type": "string",
                  "format": "date-time"
                },
                "target_period": {
                  "description": "給料明細の対象期間",
                  "type": "string"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "404": {
            "description": "NotFound",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/unprotected": {
      "get": {
        "description": "誰でも実行できるお試しAPI",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DeductionDetailsItems0": {
      "type": "object",
      "properties": {
        "amount_of_deduction_detail": {
          "description": "控除詳細の総額",
          "type": "integer",
          "format": "int32"
        },
        "nominal": {
          "description": "控除詳細の名目",
          "type": "string"
        }
      }
    },
    "EarningDetailsItems0": {
      "type": "object",
      "properties": {
        "amount_of_earning_detail": {
          "description": "支給詳細の総額",
          "type": "integer",
          "format": "int32"
        },
        "nominal": {
          "description": "支給詳細の名目",
          "type": "string"
        }
      }
    }
  }
}`))
}
