// Code generated by go-swagger; DO NOT EDIT.

package file_mgmt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"FileMgmtMicroservice/filemanagement/models"
)

// ListFilesOKCode is the HTTP code returned for type ListFilesOK
const ListFilesOKCode int = 200

/*ListFilesOK ListingFiles

swagger:response listFilesOK
*/
type ListFilesOK struct {

	/*
	  In: Body
	*/
	Payload *ListFilesOKBody `json:"body,omitempty"`
}

// NewListFilesOK creates ListFilesOK with default headers values
func NewListFilesOK() *ListFilesOK {

	return &ListFilesOK{}
}

// WithPayload adds the payload to the list files o k response
func (o *ListFilesOK) WithPayload(payload *ListFilesOKBody) *ListFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files o k response
func (o *ListFilesOK) SetPayload(payload *ListFilesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFilesInternalServerErrorCode is the HTTP code returned for type ListFilesInternalServerError
const ListFilesInternalServerErrorCode int = 500

/*ListFilesInternalServerError Server Error

swagger:response listFilesInternalServerError
*/
type ListFilesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListFilesInternalServerError creates ListFilesInternalServerError with default headers values
func NewListFilesInternalServerError() *ListFilesInternalServerError {

	return &ListFilesInternalServerError{}
}

// WithPayload adds the payload to the list files internal server error response
func (o *ListFilesInternalServerError) WithPayload(payload *models.Error) *ListFilesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files internal server error response
func (o *ListFilesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListFilesDefault error

swagger:response listFilesDefault
*/
type ListFilesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListFilesDefault creates ListFilesDefault with default headers values
func NewListFilesDefault(code int) *ListFilesDefault {
	if code <= 0 {
		code = 500
	}

	return &ListFilesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list files default response
func (o *ListFilesDefault) WithStatusCode(code int) *ListFilesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list files default response
func (o *ListFilesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list files default response
func (o *ListFilesDefault) WithPayload(payload *models.Error) *ListFilesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list files default response
func (o *ListFilesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFilesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
