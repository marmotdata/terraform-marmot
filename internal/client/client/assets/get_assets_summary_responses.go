// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/marmotdata/terraform-provider-marmot/internal/client/models"
)

// GetAssetsSummaryReader is a Reader for the GetAssetsSummary structure.
type GetAssetsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAssetsSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /assets/summary] GetAssetsSummary", response, response.Code())
	}
}

// NewGetAssetsSummaryOK creates a GetAssetsSummaryOK with default headers values
func NewGetAssetsSummaryOK() *GetAssetsSummaryOK {
	return &GetAssetsSummaryOK{}
}

/*
GetAssetsSummaryOK describes a response with status code 200, with default header values.

OK
*/
type GetAssetsSummaryOK struct {
	Payload *models.AssetsAssetSummaryResponse
}

// IsSuccess returns true when this get assets summary o k response has a 2xx status code
func (o *GetAssetsSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get assets summary o k response has a 3xx status code
func (o *GetAssetsSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets summary o k response has a 4xx status code
func (o *GetAssetsSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assets summary o k response has a 5xx status code
func (o *GetAssetsSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get assets summary o k response a status code equal to that given
func (o *GetAssetsSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get assets summary o k response
func (o *GetAssetsSummaryOK) Code() int {
	return 200
}

func (o *GetAssetsSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/summary][%d] getAssetsSummaryOK %s", 200, payload)
}

func (o *GetAssetsSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/summary][%d] getAssetsSummaryOK %s", 200, payload)
}

func (o *GetAssetsSummaryOK) GetPayload() *models.AssetsAssetSummaryResponse {
	return o.Payload
}

func (o *GetAssetsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetsAssetSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetsSummaryInternalServerError creates a GetAssetsSummaryInternalServerError with default headers values
func NewGetAssetsSummaryInternalServerError() *GetAssetsSummaryInternalServerError {
	return &GetAssetsSummaryInternalServerError{}
}

/*
GetAssetsSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAssetsSummaryInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get assets summary internal server error response has a 2xx status code
func (o *GetAssetsSummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assets summary internal server error response has a 3xx status code
func (o *GetAssetsSummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets summary internal server error response has a 4xx status code
func (o *GetAssetsSummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assets summary internal server error response has a 5xx status code
func (o *GetAssetsSummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get assets summary internal server error response a status code equal to that given
func (o *GetAssetsSummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get assets summary internal server error response
func (o *GetAssetsSummaryInternalServerError) Code() int {
	return 500
}

func (o *GetAssetsSummaryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/summary][%d] getAssetsSummaryInternalServerError %s", 500, payload)
}

func (o *GetAssetsSummaryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/summary][%d] getAssetsSummaryInternalServerError %s", 500, payload)
}

func (o *GetAssetsSummaryInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetAssetsSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
