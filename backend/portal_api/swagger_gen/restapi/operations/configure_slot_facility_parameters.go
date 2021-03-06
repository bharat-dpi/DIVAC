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
)

// NewConfigureSlotFacilityParams creates a new ConfigureSlotFacilityParams object
// no default values defined in spec.
func NewConfigureSlotFacilityParams() ConfigureSlotFacilityParams {

	return ConfigureSlotFacilityParams{}
}

// ConfigureSlotFacilityParams contains all the bound params for the configure slot facility operation
// typically these are obtained from a http.Request
//
// swagger:parameters configureSlotFacility
type ConfigureSlotFacilityParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body ConfigureSlotFacilityBody
	/*Id of facility
	  Required: true
	  In: path
	*/
	FacilityID string
	/*Id of program
	  Required: true
	  In: path
	*/
	ProgramID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewConfigureSlotFacilityParams() beforehand.
func (o *ConfigureSlotFacilityParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body ConfigureSlotFacilityBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	}
	rFacilityID, rhkFacilityID, _ := route.Params.GetOK("facilityId")
	if err := o.bindFacilityID(rFacilityID, rhkFacilityID, route.Formats); err != nil {
		res = append(res, err)
	}

	rProgramID, rhkProgramID, _ := route.Params.GetOK("programId")
	if err := o.bindProgramID(rProgramID, rhkProgramID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFacilityID binds and validates parameter FacilityID from path.
func (o *ConfigureSlotFacilityParams) bindFacilityID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.FacilityID = raw

	return nil
}

// bindProgramID binds and validates parameter ProgramID from path.
func (o *ConfigureSlotFacilityParams) bindProgramID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProgramID = raw

	return nil
}
