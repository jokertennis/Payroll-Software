package prompt

import (
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetUnprotectedHandlerStruct struct {}

// Handle executing the request and returning a response
func (s *GetUnprotectedHandlerStruct) Handle(params operations.GetUnprotectedParams) middleware.Responder {
	response := operations.NewGetUnprotectedOK().WithPayload(&operations.GetUnprotectedOKBody{
		Message: "This is the unprotected handler",
	})

	return response
}