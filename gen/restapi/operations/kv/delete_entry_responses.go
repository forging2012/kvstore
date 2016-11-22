package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/casualjim/patmosdb/gen/models"
)

/*DeleteEntryNoContent the delete was successful

swagger:response deleteEntryNoContent
*/
type DeleteEntryNoContent struct {
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`
}

// NewDeleteEntryNoContent creates DeleteEntryNoContent with default headers values
func NewDeleteEntryNoContent() *DeleteEntryNoContent {
	return &DeleteEntryNoContent{}
}

// WithXRequestID adds the xRequestId to the delete entry no content response
func (o *DeleteEntryNoContent) WithXRequestID(xRequestID string) *DeleteEntryNoContent {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete entry no content response
func (o *DeleteEntryNoContent) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *DeleteEntryNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(204)
}

/*DeleteEntryDefault Error

swagger:response deleteEntryDefault
*/
type DeleteEntryDefault struct {
	_statusCode int
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteEntryDefault creates DeleteEntryDefault with default headers values
func NewDeleteEntryDefault(code int) *DeleteEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete entry default response
func (o *DeleteEntryDefault) WithStatusCode(code int) *DeleteEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete entry default response
func (o *DeleteEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the delete entry default response
func (o *DeleteEntryDefault) WithXRequestID(xRequestID string) *DeleteEntryDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete entry default response
func (o *DeleteEntryDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete entry default response
func (o *DeleteEntryDefault) WithPayload(payload *models.Error) *DeleteEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete entry default response
func (o *DeleteEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}