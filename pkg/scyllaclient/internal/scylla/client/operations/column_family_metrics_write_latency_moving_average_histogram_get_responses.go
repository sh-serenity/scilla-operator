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

// ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGet structure.
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK creates a ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK() *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK {
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK column family metrics write latency moving average histogram get o k
*/
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK struct {
	Payload []*models.RateMovingAverageAndHistogram
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK) GetPayload() []*models.RateMovingAverageAndHistogram {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault creates a ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault(code int) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault {
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency moving average histogram get default response
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
