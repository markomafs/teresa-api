package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppDetailsParams creates a new GetAppDetailsParams object
// with the default values initialized.
func NewGetAppDetailsParams() *GetAppDetailsParams {
	var ()
	return &GetAppDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppDetailsParamsWithTimeout creates a new GetAppDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppDetailsParamsWithTimeout(timeout time.Duration) *GetAppDetailsParams {
	var ()
	return &GetAppDetailsParams{

		timeout: timeout,
	}
}

/*GetAppDetailsParams contains all the parameters to send to the API endpoint
for the get app details operation typically these are written to a http.Request
*/
type GetAppDetailsParams struct {

	/*AppID
	  App ID

	*/
	AppID int64
	/*TeamID
	  Team ID

	*/
	TeamID int64

	timeout time.Duration
}

// WithAppID adds the appID to the get app details params
func (o *GetAppDetailsParams) WithAppID(appID int64) *GetAppDetailsParams {
	o.AppID = appID
	return o
}

// WithTeamID adds the teamID to the get app details params
func (o *GetAppDetailsParams) WithTeamID(teamID int64) *GetAppDetailsParams {
	o.TeamID = teamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", swag.FormatInt64(o.AppID)); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
