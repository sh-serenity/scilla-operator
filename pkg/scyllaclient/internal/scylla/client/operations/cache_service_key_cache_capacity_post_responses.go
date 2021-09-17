// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// CacheServiceKeyCacheCapacityPostReader is a Reader for the CacheServiceKeyCacheCapacityPost structure.
type CacheServiceKeyCacheCapacityPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceKeyCacheCapacityPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceKeyCacheCapacityPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceKeyCacheCapacityPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceKeyCacheCapacityPostOK creates a CacheServiceKeyCacheCapacityPostOK with default headers values
func NewCacheServiceKeyCacheCapacityPostOK() *CacheServiceKeyCacheCapacityPostOK {
	return &CacheServiceKeyCacheCapacityPostOK{}
}

/*CacheServiceKeyCacheCapacityPostOK handles this case with default header values.

CacheServiceKeyCacheCapacityPostOK cache service key cache capacity post o k
*/
type CacheServiceKeyCacheCapacityPostOK struct {
}

func (o *CacheServiceKeyCacheCapacityPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceKeyCacheCapacityPostDefault creates a CacheServiceKeyCacheCapacityPostDefault with default headers values
func NewCacheServiceKeyCacheCapacityPostDefault(code int) *CacheServiceKeyCacheCapacityPostDefault {
	return &CacheServiceKeyCacheCapacityPostDefault{
		_statusCode: code,
	}
}

/*CacheServiceKeyCacheCapacityPostDefault handles this case with default header values.

internal server error
*/
type CacheServiceKeyCacheCapacityPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service key cache capacity post default response
func (o *CacheServiceKeyCacheCapacityPostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceKeyCacheCapacityPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceKeyCacheCapacityPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceKeyCacheCapacityPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
