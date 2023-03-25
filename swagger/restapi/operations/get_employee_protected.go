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

// GetEmployeeProtectedHandlerFunc turns a function with the right signature into a get employee protected handler
type GetEmployeeProtectedHandlerFunc func(GetEmployeeProtectedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEmployeeProtectedHandlerFunc) Handle(params GetEmployeeProtectedParams) middleware.Responder {
	return fn(params)
}

// GetEmployeeProtectedHandler interface for that can handle valid get employee protected params
type GetEmployeeProtectedHandler interface {
	Handle(GetEmployeeProtectedParams) middleware.Responder
}

// NewGetEmployeeProtected creates a new http.Handler for the get employee protected operation
func NewGetEmployeeProtected(ctx *middleware.Context, handler GetEmployeeProtectedHandler) *GetEmployeeProtected {
	return &GetEmployeeProtected{Context: ctx, Handler: handler}
}

/*
	GetEmployeeProtected swagger:route GET /employee/protected getEmployeeProtected

登録済みの従業員のみ実行可能なお試しAPI
*/
type GetEmployeeProtected struct {
	Context *middleware.Context
	Handler GetEmployeeProtectedHandler
}

func (o *GetEmployeeProtected) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetEmployeeProtectedParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetEmployeeProtectedInternalServerErrorBody get employee protected internal server error body
//
// swagger:model GetEmployeeProtectedInternalServerErrorBody
type GetEmployeeProtectedInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get employee protected internal server error body
func (o *GetEmployeeProtectedInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee protected internal server error body based on context it is used
func (o *GetEmployeeProtectedInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeProtectedInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeProtectedInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeProtectedInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeProtectedOKBody get employee protected o k body
//
// swagger:model GetEmployeeProtectedOKBody
type GetEmployeeProtectedOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get employee protected o k body
func (o *GetEmployeeProtectedOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee protected o k body based on context it is used
func (o *GetEmployeeProtectedOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeProtectedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeProtectedOKBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeProtectedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeProtectedUnauthorizedBody get employee protected unauthorized body
//
// swagger:model GetEmployeeProtectedUnauthorizedBody
type GetEmployeeProtectedUnauthorizedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get employee protected unauthorized body
func (o *GetEmployeeProtectedUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee protected unauthorized body based on context it is used
func (o *GetEmployeeProtectedUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeProtectedUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeProtectedUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeProtectedUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}