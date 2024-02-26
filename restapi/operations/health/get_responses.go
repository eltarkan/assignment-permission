// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetOKCode is the HTTP code returned for type GetOK
const GetOKCode int = 200

/*
GetOK OK

swagger:response getOK
*/
type GetOK struct {
}

// NewGetOK creates GetOK with default headers values
func NewGetOK() *GetOK {

	return &GetOK{}
}

// WriteResponse to the client
func (o *GetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetInternalServerErrorCode is the HTTP code returned for type GetInternalServerError
const GetInternalServerErrorCode int = 500

/*
GetInternalServerError Internal Server Error

swagger:response getInternalServerError
*/
type GetInternalServerError struct {
}

// NewGetInternalServerError creates GetInternalServerError with default headers values
func NewGetInternalServerError() *GetInternalServerError {

	return &GetInternalServerError{}
}

// WriteResponse to the client
func (o *GetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
