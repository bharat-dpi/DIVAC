// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteFacilityUserParams creates a new DeleteFacilityUserParams object
// no default values defined in spec.
func NewDeleteFacilityUserParams() DeleteFacilityUserParams {

	return DeleteFacilityUserParams{}
}

// DeleteFacilityUserParams contains all the bound params for the delete facility user operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteFacilityUser
type DeleteFacilityUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of facility user to delete
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteFacilityUserParams() beforehand.
func (o *DeleteFacilityUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("userId")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *DeleteFacilityUserParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.UserID = raw

	return nil
}
