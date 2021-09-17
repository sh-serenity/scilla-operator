// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigEnableCacheReader is a Reader for the FindConfigEnableCache structure.
type FindConfigEnableCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigEnableCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigEnableCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigEnableCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigEnableCacheOK creates a FindConfigEnableCacheOK with default headers values
func NewFindConfigEnableCacheOK() *FindConfigEnableCacheOK {
	return &FindConfigEnableCacheOK{}
}

/*FindConfigEnableCacheOK handles this case with default header values.

Config value
*/
type FindConfigEnableCacheOK struct {
	Payload bool
}

func (o *FindConfigEnableCacheOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigEnableCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigEnableCacheDefault creates a FindConfigEnableCacheDefault with default headers values
func NewFindConfigEnableCacheDefault(code int) *FindConfigEnableCacheDefault {
	return &FindConfigEnableCacheDefault{
		_statusCode: code,
	}
}

/*FindConfigEnableCacheDefault handles this case with default header values.

unexpected error
*/
type FindConfigEnableCacheDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config enable cache default response
func (o *FindConfigEnableCacheDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigEnableCacheDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigEnableCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigEnableCacheDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
