// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUnprotectedOKCode is the HTTP code returned for type GetUnprotectedOK
const GetUnprotectedOKCode int = 200

/*
GetUnprotectedOK OK

swagger:response getUnprotectedOK
*/
type GetUnprotectedOK struct {

	/*
	  In: Body
	*/
	Payload *GetUnprotectedOKBody `json:"body,omitempty"`
}

// NewGetUnprotectedOK creates GetUnprotectedOK with default headers values
func NewGetUnprotectedOK() *GetUnprotectedOK {

	return &GetUnprotectedOK{}
}

// WithPayload adds the payload to the get unprotected o k response
func (o *GetUnprotectedOK) WithPayload(payload *GetUnprotectedOKBody) *GetUnprotectedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unprotected o k response
func (o *GetUnprotectedOK) SetPayload(payload *GetUnprotectedOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnprotectedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}