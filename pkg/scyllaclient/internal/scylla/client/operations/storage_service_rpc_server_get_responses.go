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

// StorageServiceRPCServerGetReader is a Reader for the StorageServiceRPCServerGet structure.
type StorageServiceRPCServerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRPCServerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRPCServerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRPCServerGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRPCServerGetOK creates a StorageServiceRPCServerGetOK with default headers values
func NewStorageServiceRPCServerGetOK() *StorageServiceRPCServerGetOK {
	return &StorageServiceRPCServerGetOK{}
}

/*StorageServiceRPCServerGetOK handles this case with default header values.

StorageServiceRPCServerGetOK storage service Rpc server get o k
*/
type StorageServiceRPCServerGetOK struct {
	Payload bool
}

func (o *StorageServiceRPCServerGetOK) GetPayload() bool {
	return o.Payload
}

func (o *StorageServiceRPCServerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceRPCServerGetDefault creates a StorageServiceRPCServerGetDefault with default headers values
func NewStorageServiceRPCServerGetDefault(code int) *StorageServiceRPCServerGetDefault {
	return &StorageServiceRPCServerGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceRPCServerGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceRPCServerGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service Rpc server get default response
func (o *StorageServiceRPCServerGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRPCServerGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRPCServerGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRPCServerGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
