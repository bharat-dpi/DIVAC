// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfirmOTPHandlerFunc turns a function with the right signature into a confirm o t p handler
type ConfirmOTPHandlerFunc func(ConfirmOTPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConfirmOTPHandlerFunc) Handle(params ConfirmOTPParams) middleware.Responder {
	return fn(params)
}

// ConfirmOTPHandler interface for that can handle valid confirm o t p params
type ConfirmOTPHandler interface {
	Handle(ConfirmOTPParams) middleware.Responder
}

// NewConfirmOTP creates a new http.Handler for the confirm o t p operation
func NewConfirmOTP(ctx *middleware.Context, handler ConfirmOTPHandler) *ConfirmOTP {
	return &ConfirmOTP{Context: ctx, Handler: handler}
}

/*ConfirmOTP swagger:route POST /auth/confirm confirmOTP

Confirm OTP

*/
type ConfirmOTP struct {
	Context *middleware.Context
	Handler ConfirmOTPHandler
}

func (o *ConfirmOTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConfirmOTPParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ConfirmOTPBody confirm o t p body
//
// swagger:model ConfirmOTPBody
type ConfirmOTPBody struct {

	// mobile
	// Required: true
	Mobile *string `json:"mobile"`

	// otp
	// Required: true
	Otp *string `json:"otp"`

	// txn
	// Required: true
	Txn *string `json:"txn"`
}

// Validate validates this confirm o t p body
func (o *ConfirmOTPBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOtp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTxn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConfirmOTPBody) validateMobile(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"mobile", "body", o.Mobile); err != nil {
		return err
	}

	return nil
}

func (o *ConfirmOTPBody) validateOtp(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"otp", "body", o.Otp); err != nil {
		return err
	}

	return nil
}

func (o *ConfirmOTPBody) validateTxn(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"txn", "body", o.Txn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConfirmOTPBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfirmOTPBody) UnmarshalBinary(b []byte) error {
	var res ConfirmOTPBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ConfirmOTPOKBody confirm o t p o k body
//
// swagger:model ConfirmOTPOKBody
type ConfirmOTPOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this confirm o t p o k body
func (o *ConfirmOTPOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConfirmOTPOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfirmOTPOKBody) UnmarshalBinary(b []byte) error {
	var res ConfirmOTPOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
