// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteStorageFromGroupHandlerFunc turns a function with the right signature into a delete storage from group handler
type DeleteStorageFromGroupHandlerFunc func(DeleteStorageFromGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageFromGroupHandlerFunc) Handle(params DeleteStorageFromGroupParams) middleware.Responder {
	return fn(params)
}

// DeleteStorageFromGroupHandler interface for that can handle valid delete storage from group params
type DeleteStorageFromGroupHandler interface {
	Handle(DeleteStorageFromGroupParams) middleware.Responder
}

// NewDeleteStorageFromGroup creates a new http.Handler for the delete storage from group operation
func NewDeleteStorageFromGroup(ctx *middleware.Context, handler DeleteStorageFromGroupHandler) *DeleteStorageFromGroup {
	return &DeleteStorageFromGroup{Context: ctx, Handler: handler}
}

/*DeleteStorageFromGroup swagger:route DELETE /cc/v1/groups/groupStorage/id/{id} Groups deleteStorageFromGroup

Returns a GroupStorage.

Optional extended description in Markdown.

*/
type DeleteStorageFromGroup struct {
	Context *middleware.Context
	Handler DeleteStorageFromGroupHandler
}

func (o *DeleteStorageFromGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageFromGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}