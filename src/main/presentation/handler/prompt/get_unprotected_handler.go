package prompt

import (
	"github.com/jokertennis/Payroll-Software/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetUnprotectedHandlerStruct struct {}

func (s *GetUnprotectedHandlerStruct) Handle(params operations.GetUnprotectedParams) middleware.Responder {
	response := operations.NewGetUnprotectedOK().WithPayload(&operations.GetUnprotectedOKBody{
		Message: "This is the unprotected handler",
	})

	return response
}