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

// GetAssetsListReader is a Reader for the GetAssetsList structure.
type GetAssetsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAssetsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /assets/list] GetAssetsList", response, response.Code())
	}
}

// NewGetAssetsListOK creates a GetAssetsListOK with default headers values
func NewGetAssetsListOK() *GetAssetsListOK {
	return &GetAssetsListOK{}
}

/*
GetAssetsListOK describes a response with status code 200, with default header values.

OK
*/
type GetAssetsListOK struct {
	Payload *models.AssetListResult
}

// IsSuccess returns true when this get assets list o k response has a 2xx status code
func (o *GetAssetsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get assets list o k response has a 3xx status code
func (o *GetAssetsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets list o k response has a 4xx status code
func (o *GetAssetsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assets list o k response has a 5xx status code
func (o *GetAssetsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get assets list o k response a status code equal to that given
func (o *GetAssetsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get assets list o k response
func (o *GetAssetsListOK) Code() int {
	return 200
}

func (o *GetAssetsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/list][%d] getAssetsListOK %s", 200, payload)
}

func (o *GetAssetsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/list][%d] getAssetsListOK %s", 200, payload)
}

func (o *GetAssetsListOK) GetPayload() *models.AssetListResult {
	return o.Payload
}

func (o *GetAssetsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetsListInternalServerError creates a GetAssetsListInternalServerError with default headers values
func NewGetAssetsListInternalServerError() *GetAssetsListInternalServerError {
	return &GetAssetsListInternalServerError{}
}

/*
GetAssetsListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAssetsListInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get assets list internal server error response has a 2xx status code
func (o *GetAssetsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assets list internal server error response has a 3xx status code
func (o *GetAssetsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets list internal server error response has a 4xx status code
func (o *GetAssetsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assets list internal server error response has a 5xx status code
func (o *GetAssetsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get assets list internal server error response a status code equal to that given
func (o *GetAssetsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get assets list internal server error response
func (o *GetAssetsListInternalServerError) Code() int {
	return 500
}

func (o *GetAssetsListInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/list][%d] getAssetsListInternalServerError %s", 500, payload)
}

func (o *GetAssetsListInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/list][%d] getAssetsListInternalServerError %s", 500, payload)
}

func (o *GetAssetsListInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetAssetsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
