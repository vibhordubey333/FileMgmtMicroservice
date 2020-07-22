// Code generated by go-swagger; DO NOT EDIT.

package file_mgmt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFileHandlerFunc turns a function with the right signature into a get file handler
type GetFileHandlerFunc func(GetFileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFileHandlerFunc) Handle(params GetFileParams) middleware.Responder {
	return fn(params)
}

// GetFileHandler interface for that can handle valid get file params
type GetFileHandler interface {
	Handle(GetFileParams) middleware.Responder
}

// NewGetFile creates a new http.Handler for the get file operation
func NewGetFile(ctx *middleware.Context, handler GetFileHandler) *GetFile {
	return &GetFile{Context: ctx, Handler: handler}
}

/*GetFile swagger:route GET /fileMgmt/getFile fileMgmt getFile

Get a file from storage

Get a file from storage

*/
type GetFile struct {
	Context *middleware.Context
	Handler GetFileHandler
}

func (o *GetFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFileParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}