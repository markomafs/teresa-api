package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*GetTeamsOK Get teams

swagger:response getTeamsOK
*/
type GetTeamsOK struct {

	// In: body
	Payload GetTeamsOKBodyBody `json:"body,omitempty"`
}

// NewGetTeamsOK creates GetTeamsOK with default headers values
func NewGetTeamsOK() *GetTeamsOK {
	return &GetTeamsOK{}
}

// WithPayload adds the payload to the get teams o k response
func (o *GetTeamsOK) WithPayload(payload GetTeamsOKBodyBody) *GetTeamsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams o k response
func (o *GetTeamsOK) SetPayload(payload GetTeamsOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetTeamsBadRequest User sent a bad request

swagger:response getTeamsBadRequest
*/
type GetTeamsBadRequest struct {

	// In: body
	Payload *models.BadRequest `json:"body,omitempty"`
}

// NewGetTeamsBadRequest creates GetTeamsBadRequest with default headers values
func NewGetTeamsBadRequest() *GetTeamsBadRequest {
	return &GetTeamsBadRequest{}
}

// WithPayload adds the payload to the get teams bad request response
func (o *GetTeamsBadRequest) WithPayload(payload *models.BadRequest) *GetTeamsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams bad request response
func (o *GetTeamsBadRequest) SetPayload(payload *models.BadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTeamsUnauthorized User not authorized

swagger:response getTeamsUnauthorized
*/
type GetTeamsUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetTeamsUnauthorized creates GetTeamsUnauthorized with default headers values
func NewGetTeamsUnauthorized() *GetTeamsUnauthorized {
	return &GetTeamsUnauthorized{}
}

// WithPayload adds the payload to the get teams unauthorized response
func (o *GetTeamsUnauthorized) WithPayload(payload *models.Unauthorized) *GetTeamsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams unauthorized response
func (o *GetTeamsUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTeamsForbidden User does not have the credentials to access this resource


swagger:response getTeamsForbidden
*/
type GetTeamsForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetTeamsForbidden creates GetTeamsForbidden with default headers values
func NewGetTeamsForbidden() *GetTeamsForbidden {
	return &GetTeamsForbidden{}
}

// WithPayload adds the payload to the get teams forbidden response
func (o *GetTeamsForbidden) WithPayload(payload *models.Unauthorized) *GetTeamsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams forbidden response
func (o *GetTeamsForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTeamsNotFound Resource not found

swagger:response getTeamsNotFound
*/
type GetTeamsNotFound struct {

	// In: body
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetTeamsNotFound creates GetTeamsNotFound with default headers values
func NewGetTeamsNotFound() *GetTeamsNotFound {
	return &GetTeamsNotFound{}
}

// WithPayload adds the payload to the get teams not found response
func (o *GetTeamsNotFound) WithPayload(payload *models.NotFound) *GetTeamsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams not found response
func (o *GetTeamsNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTeamsDefault User not authorized

swagger:response getTeamsDefault
*/
type GetTeamsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetTeamsDefault creates GetTeamsDefault with default headers values
func NewGetTeamsDefault(code int) *GetTeamsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTeamsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get teams default response
func (o *GetTeamsDefault) WithStatusCode(code int) *GetTeamsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get teams default response
func (o *GetTeamsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get teams default response
func (o *GetTeamsDefault) WithPayload(payload *models.Unauthorized) *GetTeamsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams default response
func (o *GetTeamsDefault) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
