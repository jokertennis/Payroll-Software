// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetEmployeeSalaryStatementParams creates a new GetEmployeeSalaryStatementParams object
//
// There are no default values defined in the spec.
func NewGetEmployeeSalaryStatementParams() GetEmployeeSalaryStatementParams {

	return GetEmployeeSalaryStatementParams{}
}

// GetEmployeeSalaryStatementParams contains all the bound params for the get employee salary statement operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetEmployeeSalaryStatement
type GetEmployeeSalaryStatementParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*欲しい給料明細の月
	  Required: true
	  Maximum: 12
	  Minimum: 1
	  In: query
	*/
	Month int32
	/*欲しい給料明細の年
	  Required: true
	  Maximum: 3000
	  Minimum: 1500
	  In: query
	*/
	Year int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEmployeeSalaryStatementParams() beforehand.
func (o *GetEmployeeSalaryStatementParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMonth, qhkMonth, _ := qs.GetOK("month")
	if err := o.bindMonth(qMonth, qhkMonth, route.Formats); err != nil {
		res = append(res, err)
	}

	qYear, qhkYear, _ := qs.GetOK("year")
	if err := o.bindYear(qYear, qhkYear, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMonth binds and validates parameter Month from query.
func (o *GetEmployeeSalaryStatementParams) bindMonth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("month", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("month", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("month", "query", "int32", raw)
	}
	o.Month = value

	if err := o.validateMonth(formats); err != nil {
		return err
	}

	return nil
}

// validateMonth carries on validations for parameter Month
func (o *GetEmployeeSalaryStatementParams) validateMonth(formats strfmt.Registry) error {

	if err := validate.MinimumInt("month", "query", int64(o.Month), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("month", "query", int64(o.Month), 12, false); err != nil {
		return err
	}

	return nil
}

// bindYear binds and validates parameter Year from query.
func (o *GetEmployeeSalaryStatementParams) bindYear(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("year", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("year", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("year", "query", "int64", raw)
	}
	o.Year = value

	if err := o.validateYear(formats); err != nil {
		return err
	}

	return nil
}

// validateYear carries on validations for parameter Year
func (o *GetEmployeeSalaryStatementParams) validateYear(formats strfmt.Registry) error {

	if err := validate.MinimumInt("year", "query", o.Year, 1500, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("year", "query", o.Year, 3000, false); err != nil {
		return err
	}

	return nil
}
