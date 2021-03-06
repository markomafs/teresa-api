package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTeamDetailHandlerFunc turns a function with the right signature into a get team detail handler
type GetTeamDetailHandlerFunc func(GetTeamDetailParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeamDetailHandlerFunc) Handle(params GetTeamDetailParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTeamDetailHandler interface for that can handle valid get team detail params
type GetTeamDetailHandler interface {
	Handle(GetTeamDetailParams, interface{}) middleware.Responder
}

// NewGetTeamDetail creates a new http.Handler for the get team detail operation
func NewGetTeamDetail(ctx *middleware.Context, handler GetTeamDetailHandler) *GetTeamDetail {
	return &GetTeamDetail{Context: ctx, Handler: handler}
}

/*GetTeamDetail swagger:route GET /teams/{team_id} teams getTeamDetail

Get team details

Get the details of a specific team

*/
type GetTeamDetail struct {
	Context *middleware.Context
	Handler GetTeamDetailHandler
}

func (o *GetTeamDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTeamDetailParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
