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

// DeleteAppointmentHandlerFunc turns a function with the right signature into a delete appointment handler
type DeleteAppointmentHandlerFunc func(DeleteAppointmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAppointmentHandlerFunc) Handle(params DeleteAppointmentParams) middleware.Responder {
	return fn(params)
}

// DeleteAppointmentHandler interface for that can handle valid delete appointment params
type DeleteAppointmentHandler interface {
	Handle(DeleteAppointmentParams) middleware.Responder
}

// NewDeleteAppointment creates a new http.Handler for the delete appointment operation
func NewDeleteAppointment(ctx *middleware.Context, handler DeleteAppointmentHandler) *DeleteAppointment {
	return &DeleteAppointment{Context: ctx, Handler: handler}
}

/*DeleteAppointment swagger:route DELETE /appointment deleteAppointment

Delete the appointment

*/
type DeleteAppointment struct {
	Context *middleware.Context
	Handler DeleteAppointmentHandler
}

func (o *DeleteAppointment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAppointmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteAppointmentBadRequestBody delete appointment bad request body
//
// swagger:model DeleteAppointmentBadRequestBody
type DeleteAppointmentBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete appointment bad request body
func (o *DeleteAppointmentBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAppointmentBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAppointmentBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteAppointmentBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteAppointmentBody delete appointment body
//
// swagger:model DeleteAppointmentBody
type DeleteAppointmentBody struct {

	// enrollment code
	// Required: true
	EnrollmentCode *string `json:"enrollmentCode"`
}

// Validate validates this delete appointment body
func (o *DeleteAppointmentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnrollmentCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAppointmentBody) validateEnrollmentCode(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"enrollmentCode", "body", o.EnrollmentCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAppointmentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAppointmentBody) UnmarshalBinary(b []byte) error {
	var res DeleteAppointmentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
