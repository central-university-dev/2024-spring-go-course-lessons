// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateTaskBadRequestCode is the HTTP code returned for type CreateTaskBadRequest
const CreateTaskBadRequestCode int = 400

/*
CreateTaskBadRequest Некорректный ввод

swagger:response createTaskBadRequest
*/
type CreateTaskBadRequest struct {
}

// NewCreateTaskBadRequest creates CreateTaskBadRequest with default headers values
func NewCreateTaskBadRequest() *CreateTaskBadRequest {

	return &CreateTaskBadRequest{}
}

// WriteResponse to the client
func (o *CreateTaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateTaskUnprocessableEntityCode is the HTTP code returned for type CreateTaskUnprocessableEntity
const CreateTaskUnprocessableEntityCode int = 422

/*
CreateTaskUnprocessableEntity Недопустимые значения

swagger:response createTaskUnprocessableEntity
*/
type CreateTaskUnprocessableEntity struct {
}

// NewCreateTaskUnprocessableEntity creates CreateTaskUnprocessableEntity with default headers values
func NewCreateTaskUnprocessableEntity() *CreateTaskUnprocessableEntity {

	return &CreateTaskUnprocessableEntity{}
}

// WriteResponse to the client
func (o *CreateTaskUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(422)
}