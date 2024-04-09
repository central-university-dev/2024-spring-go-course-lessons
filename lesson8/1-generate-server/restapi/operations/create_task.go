// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateTaskHandlerFunc turns a function with the right signature into a create task handler
type CreateTaskHandlerFunc func(CreateTaskParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTaskHandlerFunc) Handle(params CreateTaskParams) middleware.Responder {
	return fn(params)
}

// CreateTaskHandler interface for that can handle valid create task params
type CreateTaskHandler interface {
	Handle(CreateTaskParams) middleware.Responder
}

// NewCreateTask creates a new http.Handler for the create task operation
func NewCreateTask(ctx *middleware.Context, handler CreateTaskHandler) *CreateTask {
	return &CreateTask{Context: ctx, Handler: handler}
}

/*
	CreateTask swagger:route POST /task createTask

Добавить новую задачу
*/
type CreateTask struct {
	Context *middleware.Context
	Handler CreateTaskHandler
}

func (o *CreateTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateTaskParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}