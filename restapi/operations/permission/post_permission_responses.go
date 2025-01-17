// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"assignment-permission/models"
)

// PostPermissionOKCode is the HTTP code returned for type PostPermissionOK
const PostPermissionOKCode int = 200

/*
PostPermissionOK OK

swagger:response postPermissionOK
*/
type PostPermissionOK struct {

	/*
	  In: Body
	*/
	Payload *models.PermissionCreateSuccessResponse `json:"body,omitempty"`
}

// NewPostPermissionOK creates PostPermissionOK with default headers values
func NewPostPermissionOK() *PostPermissionOK {

	return &PostPermissionOK{}
}

// WithPayload adds the payload to the post permission o k response
func (o *PostPermissionOK) WithPayload(payload *models.PermissionCreateSuccessResponse) *PostPermissionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post permission o k response
func (o *PostPermissionOK) SetPayload(payload *models.PermissionCreateSuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPermissionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPermissionInternalServerErrorCode is the HTTP code returned for type PostPermissionInternalServerError
const PostPermissionInternalServerErrorCode int = 500

/*
PostPermissionInternalServerError Internal Server Error

swagger:response postPermissionInternalServerError
*/
type PostPermissionInternalServerError struct {
}

// NewPostPermissionInternalServerError creates PostPermissionInternalServerError with default headers values
func NewPostPermissionInternalServerError() *PostPermissionInternalServerError {

	return &PostPermissionInternalServerError{}
}

// WriteResponse to the client
func (o *PostPermissionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
