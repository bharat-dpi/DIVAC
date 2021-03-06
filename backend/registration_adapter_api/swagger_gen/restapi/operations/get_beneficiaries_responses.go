// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetBeneficiariesOKCode is the HTTP code returned for type GetBeneficiariesOK
const GetBeneficiariesOKCode int = 200

/*GetBeneficiariesOK OK

swagger:response getBeneficiariesOK
*/
type GetBeneficiariesOK struct {

	/*
	  In: Body
	*/
	Payload []*GetBeneficiariesOKBodyItems0 `json:"body,omitempty"`
}

// NewGetBeneficiariesOK creates GetBeneficiariesOK with default headers values
func NewGetBeneficiariesOK() *GetBeneficiariesOK {

	return &GetBeneficiariesOK{}
}

// WithPayload adds the payload to the get beneficiaries o k response
func (o *GetBeneficiariesOK) WithPayload(payload []*GetBeneficiariesOKBodyItems0) *GetBeneficiariesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get beneficiaries o k response
func (o *GetBeneficiariesOK) SetPayload(payload []*GetBeneficiariesOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBeneficiariesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetBeneficiariesOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBeneficiariesBadRequestCode is the HTTP code returned for type GetBeneficiariesBadRequest
const GetBeneficiariesBadRequestCode int = 400

/*GetBeneficiariesBadRequest Bad Request

swagger:response getBeneficiariesBadRequest
*/
type GetBeneficiariesBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetBeneficiariesBadRequest creates GetBeneficiariesBadRequest with default headers values
func NewGetBeneficiariesBadRequest() *GetBeneficiariesBadRequest {

	return &GetBeneficiariesBadRequest{}
}

// WithPayload adds the payload to the get beneficiaries bad request response
func (o *GetBeneficiariesBadRequest) WithPayload(payload interface{}) *GetBeneficiariesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get beneficiaries bad request response
func (o *GetBeneficiariesBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBeneficiariesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBeneficiariesUnauthorizedCode is the HTTP code returned for type GetBeneficiariesUnauthorized
const GetBeneficiariesUnauthorizedCode int = 401

/*GetBeneficiariesUnauthorized Unauthorized

swagger:response getBeneficiariesUnauthorized
*/
type GetBeneficiariesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetBeneficiariesUnauthorized creates GetBeneficiariesUnauthorized with default headers values
func NewGetBeneficiariesUnauthorized() *GetBeneficiariesUnauthorized {

	return &GetBeneficiariesUnauthorized{}
}

// WithPayload adds the payload to the get beneficiaries unauthorized response
func (o *GetBeneficiariesUnauthorized) WithPayload(payload interface{}) *GetBeneficiariesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get beneficiaries unauthorized response
func (o *GetBeneficiariesUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBeneficiariesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
