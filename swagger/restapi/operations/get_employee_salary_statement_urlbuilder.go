// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetEmployeeSalaryStatementURL generates an URL for the get employee salary statement operation
type GetEmployeeSalaryStatementURL struct {
	Month int32
	Year  int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetEmployeeSalaryStatementURL) WithBasePath(bp string) *GetEmployeeSalaryStatementURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetEmployeeSalaryStatementURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetEmployeeSalaryStatementURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/employee/salary_statement"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	monthQ := swag.FormatInt32(o.Month)
	if monthQ != "" {
		qs.Set("month", monthQ)
	}

	yearQ := swag.FormatInt32(o.Year)
	if yearQ != "" {
		qs.Set("year", yearQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetEmployeeSalaryStatementURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetEmployeeSalaryStatementURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetEmployeeSalaryStatementURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetEmployeeSalaryStatementURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetEmployeeSalaryStatementURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetEmployeeSalaryStatementURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}