// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetEmployeeSalaryStatementHandlerFunc turns a function with the right signature into a get employee salary statement handler
type GetEmployeeSalaryStatementHandlerFunc func(GetEmployeeSalaryStatementParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEmployeeSalaryStatementHandlerFunc) Handle(params GetEmployeeSalaryStatementParams) middleware.Responder {
	return fn(params)
}

// GetEmployeeSalaryStatementHandler interface for that can handle valid get employee salary statement params
type GetEmployeeSalaryStatementHandler interface {
	Handle(GetEmployeeSalaryStatementParams) middleware.Responder
}

// NewGetEmployeeSalaryStatement creates a new http.Handler for the get employee salary statement operation
func NewGetEmployeeSalaryStatement(ctx *middleware.Context, handler GetEmployeeSalaryStatementHandler) *GetEmployeeSalaryStatement {
	return &GetEmployeeSalaryStatement{Context: ctx, Handler: handler}
}

/*
	GetEmployeeSalaryStatement swagger:route GET /employee/salary_statement getEmployeeSalaryStatement

特定の年月の給料明細を取得するための従業員向けAPI
*/
type GetEmployeeSalaryStatement struct {
	Context *middleware.Context
	Handler GetEmployeeSalaryStatementHandler
}

func (o *GetEmployeeSalaryStatement) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetEmployeeSalaryStatementParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetEmployeeSalaryStatementInternalServerErrorBody get employee salary statement internal server error body
//
// swagger:model GetEmployeeSalaryStatementInternalServerErrorBody
type GetEmployeeSalaryStatementInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get employee salary statement internal server error body
func (o *GetEmployeeSalaryStatementInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee salary statement internal server error body based on context it is used
func (o *GetEmployeeSalaryStatementInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeSalaryStatementInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeSalaryStatementNotFoundBody get employee salary statement not found body
//
// swagger:model GetEmployeeSalaryStatementNotFoundBody
type GetEmployeeSalaryStatementNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get employee salary statement not found body
func (o *GetEmployeeSalaryStatementNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee salary statement not found body based on context it is used
func (o *GetEmployeeSalaryStatementNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeSalaryStatementNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeSalaryStatementOKBody get employee salary statement o k body
//
// swagger:model GetEmployeeSalaryStatementOKBody
type GetEmployeeSalaryStatementOKBody struct {

	// 控除の総額
	AmountOfDeduction int32 `json:"amount_of_deduction,omitempty"`

	// 支給の総額
	AmountOfEarning int32 `json:"amount_of_earning,omitempty"`

	// deduction details
	DeductionDetails []*GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0 `json:"deduction_details"`

	// earning details
	EarningDetails []*GetEmployeeSalaryStatementOKBodyEarningDetailsItems0 `json:"earning_details"`

	// 従業員名
	NameOfEmployee string `json:"name_of_employee,omitempty"`

	// 給料支払い日時
	// Format: date-time
	Payday strfmt.DateTime `json:"payday,omitempty"`

	// 給料明細の対象期間
	TargetPeriod string `json:"target_period,omitempty"`
}

// Validate validates this get employee salary statement o k body
func (o *GetEmployeeSalaryStatementOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeductionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEarningDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePayday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEmployeeSalaryStatementOKBody) validateDeductionDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.DeductionDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.DeductionDetails); i++ {
		if swag.IsZero(o.DeductionDetails[i]) { // not required
			continue
		}

		if o.DeductionDetails[i] != nil {
			if err := o.DeductionDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEmployeeSalaryStatementOK" + "." + "deduction_details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEmployeeSalaryStatementOK" + "." + "deduction_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetEmployeeSalaryStatementOKBody) validateEarningDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.EarningDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.EarningDetails); i++ {
		if swag.IsZero(o.EarningDetails[i]) { // not required
			continue
		}

		if o.EarningDetails[i] != nil {
			if err := o.EarningDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEmployeeSalaryStatementOK" + "." + "earning_details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEmployeeSalaryStatementOK" + "." + "earning_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetEmployeeSalaryStatementOKBody) validatePayday(formats strfmt.Registry) error {
	if swag.IsZero(o.Payday) { // not required
		return nil
	}

	if err := validate.FormatOf("getEmployeeSalaryStatementOK"+"."+"payday", "body", "date-time", o.Payday.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get employee salary statement o k body based on the context it is used
func (o *GetEmployeeSalaryStatementOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDeductionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateEarningDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEmployeeSalaryStatementOKBody) contextValidateDeductionDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DeductionDetails); i++ {

		if o.DeductionDetails[i] != nil {
			if err := o.DeductionDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEmployeeSalaryStatementOK" + "." + "deduction_details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEmployeeSalaryStatementOK" + "." + "deduction_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetEmployeeSalaryStatementOKBody) contextValidateEarningDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.EarningDetails); i++ {

		if o.EarningDetails[i] != nil {
			if err := o.EarningDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEmployeeSalaryStatementOK" + "." + "earning_details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEmployeeSalaryStatementOK" + "." + "earning_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementOKBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeSalaryStatementOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0 get employee salary statement o k body deduction details items0
//
// swagger:model GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0
type GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0 struct {

	// 控除詳細の総額
	AmountOfDeductionDetail int32 `json:"amount_of_deduction_detail,omitempty"`

	// 控除詳細の名目
	Nominal string `json:"nominal,omitempty"`
}

// Validate validates this get employee salary statement o k body deduction details items0
func (o *GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee salary statement o k body deduction details items0 based on context it is used
func (o *GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeSalaryStatementOKBodyEarningDetailsItems0 get employee salary statement o k body earning details items0
//
// swagger:model GetEmployeeSalaryStatementOKBodyEarningDetailsItems0
type GetEmployeeSalaryStatementOKBodyEarningDetailsItems0 struct {

	// 支給詳細の総額
	AmountOfEarningDetail int32 `json:"amount_of_earning_detail,omitempty"`

	// 支給詳細の名目
	Nominal string `json:"nominal,omitempty"`
}

// Validate validates this get employee salary statement o k body earning details items0
func (o *GetEmployeeSalaryStatementOKBodyEarningDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee salary statement o k body earning details items0 based on context it is used
func (o *GetEmployeeSalaryStatementOKBodyEarningDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementOKBodyEarningDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementOKBodyEarningDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetEmployeeSalaryStatementOKBodyEarningDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEmployeeSalaryStatementUnauthorizedBody get employee salary statement unauthorized body
//
// swagger:model GetEmployeeSalaryStatementUnauthorizedBody
type GetEmployeeSalaryStatementUnauthorizedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get employee salary statement unauthorized body
func (o *GetEmployeeSalaryStatementUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get employee salary statement unauthorized body based on context it is used
func (o *GetEmployeeSalaryStatementUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeeSalaryStatementUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeeSalaryStatementUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
