// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostAdministratorSalaryStatementIndividualHandlerFunc turns a function with the right signature into a post administrator salary statement individual handler
type PostAdministratorSalaryStatementIndividualHandlerFunc func(PostAdministratorSalaryStatementIndividualParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAdministratorSalaryStatementIndividualHandlerFunc) Handle(params PostAdministratorSalaryStatementIndividualParams) middleware.Responder {
	return fn(params)
}

// PostAdministratorSalaryStatementIndividualHandler interface for that can handle valid post administrator salary statement individual params
type PostAdministratorSalaryStatementIndividualHandler interface {
	Handle(PostAdministratorSalaryStatementIndividualParams) middleware.Responder
}

// NewPostAdministratorSalaryStatementIndividual creates a new http.Handler for the post administrator salary statement individual operation
func NewPostAdministratorSalaryStatementIndividual(ctx *middleware.Context, handler PostAdministratorSalaryStatementIndividualHandler) *PostAdministratorSalaryStatementIndividual {
	return &PostAdministratorSalaryStatementIndividual{Context: ctx, Handler: handler}
}

/*
	PostAdministratorSalaryStatementIndividual swagger:route POST /administrator/salary_statement_individual postAdministratorSalaryStatementIndividual

特定の年と月の給料明細を個別支給・個別控除を利用して作成するための管理者向けAPI
*/
type PostAdministratorSalaryStatementIndividual struct {
	Context *middleware.Context
	Handler PostAdministratorSalaryStatementIndividualHandler
}

func (o *PostAdministratorSalaryStatementIndividual) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAdministratorSalaryStatementIndividualParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAdministratorSalaryStatementIndividualBadRequestBody post administrator salary statement individual bad request body
//
// swagger:model PostAdministratorSalaryStatementIndividualBadRequestBody
type PostAdministratorSalaryStatementIndividualBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post administrator salary statement individual bad request body
func (o *PostAdministratorSalaryStatementIndividualBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post administrator salary statement individual bad request body based on context it is used
func (o *PostAdministratorSalaryStatementIndividualBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostAdministratorSalaryStatementIndividualBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdministratorSalaryStatementIndividualCreatedBody post administrator salary statement individual created body
//
// swagger:model PostAdministratorSalaryStatementIndividualCreatedBody
type PostAdministratorSalaryStatementIndividualCreatedBody struct {

	// 給料明細id
	IDOfSalaryStatement int64 `json:"id_of_salary_statement,omitempty"`
}

// Validate validates this post administrator salary statement individual created body
func (o *PostAdministratorSalaryStatementIndividualCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post administrator salary statement individual created body based on context it is used
func (o *PostAdministratorSalaryStatementIndividualCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostAdministratorSalaryStatementIndividualCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdministratorSalaryStatementIndividualInternalServerErrorBody post administrator salary statement individual internal server error body
//
// swagger:model PostAdministratorSalaryStatementIndividualInternalServerErrorBody
type PostAdministratorSalaryStatementIndividualInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post administrator salary statement individual internal server error body
func (o *PostAdministratorSalaryStatementIndividualInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post administrator salary statement individual internal server error body based on context it is used
func (o *PostAdministratorSalaryStatementIndividualInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostAdministratorSalaryStatementIndividualInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAdministratorSalaryStatementIndividualUnauthorizedBody post administrator salary statement individual unauthorized body
//
// swagger:model PostAdministratorSalaryStatementIndividualUnauthorizedBody
type PostAdministratorSalaryStatementIndividualUnauthorizedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post administrator salary statement individual unauthorized body
func (o *PostAdministratorSalaryStatementIndividualUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post administrator salary statement individual unauthorized body based on context it is used
func (o *PostAdministratorSalaryStatementIndividualUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAdministratorSalaryStatementIndividualUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostAdministratorSalaryStatementIndividualUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
