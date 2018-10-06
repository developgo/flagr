// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/checkr/flagr/swagger_gen/models"
)

// GetExportSidecarOKCode is the HTTP code returned for type GetExportSidecarOK
const GetExportSidecarOKCode int = 200

/*GetExportSidecarOK OK

swagger:response getExportSidecarOK
*/
type GetExportSidecarOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetExportSidecarOK creates GetExportSidecarOK with default headers values
func NewGetExportSidecarOK() *GetExportSidecarOK {

	return &GetExportSidecarOK{}
}

// WithPayload adds the payload to the get export sidecar o k response
func (o *GetExportSidecarOK) WithPayload(payload io.ReadCloser) *GetExportSidecarOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get export sidecar o k response
func (o *GetExportSidecarOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExportSidecarOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetExportSidecarDefault generic error response

swagger:response getExportSidecarDefault
*/
type GetExportSidecarDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetExportSidecarDefault creates GetExportSidecarDefault with default headers values
func NewGetExportSidecarDefault(code int) *GetExportSidecarDefault {
	if code <= 0 {
		code = 500
	}

	return &GetExportSidecarDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get export sidecar default response
func (o *GetExportSidecarDefault) WithStatusCode(code int) *GetExportSidecarDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get export sidecar default response
func (o *GetExportSidecarDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get export sidecar default response
func (o *GetExportSidecarDefault) WithPayload(payload *models.Error) *GetExportSidecarDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get export sidecar default response
func (o *GetExportSidecarDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExportSidecarDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}