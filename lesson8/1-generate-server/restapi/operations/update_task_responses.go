// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateTaskBadRequestCode is the HTTP code returned for type UpdateTaskBadRequest
const UpdateTaskBadRequestCode int = 400

/*
UpdateTaskBadRequest Некорректный ввод

swagger:response updateTaskBadRequest
*/
type UpdateTaskBadRequest struct {
}

// NewUpdateTaskBadRequest creates UpdateTaskBadRequest with default headers values
func NewUpdateTaskBadRequest() *UpdateTaskBadRequest {

	return &UpdateTaskBadRequest{}
}

// WriteResponse to the client
func (o *UpdateTaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateTaskNotFoundCode is the HTTP code returned for type UpdateTaskNotFound
const UpdateTaskNotFoundCode int = 404

/*
UpdateTaskNotFound Такой задачи нет

swagger:response updateTaskNotFound
*/
type UpdateTaskNotFound struct {
}

// NewUpdateTaskNotFound creates UpdateTaskNotFound with default headers values
func NewUpdateTaskNotFound() *UpdateTaskNotFound {

	return &UpdateTaskNotFound{}
}

// WriteResponse to the client
func (o *UpdateTaskNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateTaskUnprocessableEntityCode is the HTTP code returned for type UpdateTaskUnprocessableEntity
const UpdateTaskUnprocessableEntityCode int = 422

/*
UpdateTaskUnprocessableEntity Недопустимые значения

swagger:response updateTaskUnprocessableEntity
*/
type UpdateTaskUnprocessableEntity struct {
}

// NewUpdateTaskUnprocessableEntity creates UpdateTaskUnprocessableEntity with default headers values
func NewUpdateTaskUnprocessableEntity() *UpdateTaskUnprocessableEntity {

	return &UpdateTaskUnprocessableEntity{}
}

// WriteResponse to the client
func (o *UpdateTaskUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(422)
}