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

// GetAssetsDocumentationMrnReader is a Reader for the GetAssetsDocumentationMrn structure.
type GetAssetsDocumentationMrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetsDocumentationMrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetsDocumentationMrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAssetsDocumentationMrnNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAssetsDocumentationMrnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /assets/documentation/{mrn}] GetAssetsDocumentationMrn", response, response.Code())
	}
}

// NewGetAssetsDocumentationMrnOK creates a GetAssetsDocumentationMrnOK with default headers values
func NewGetAssetsDocumentationMrnOK() *GetAssetsDocumentationMrnOK {
	return &GetAssetsDocumentationMrnOK{}
}

/*
GetAssetsDocumentationMrnOK describes a response with status code 200, with default header values.

OK
*/
type GetAssetsDocumentationMrnOK struct {
	Payload []*models.AssetdocsDocumentation
}

// IsSuccess returns true when this get assets documentation mrn o k response has a 2xx status code
func (o *GetAssetsDocumentationMrnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get assets documentation mrn o k response has a 3xx status code
func (o *GetAssetsDocumentationMrnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets documentation mrn o k response has a 4xx status code
func (o *GetAssetsDocumentationMrnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assets documentation mrn o k response has a 5xx status code
func (o *GetAssetsDocumentationMrnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get assets documentation mrn o k response a status code equal to that given
func (o *GetAssetsDocumentationMrnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get assets documentation mrn o k response
func (o *GetAssetsDocumentationMrnOK) Code() int {
	return 200
}

func (o *GetAssetsDocumentationMrnOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/documentation/{mrn}][%d] getAssetsDocumentationMrnOK %s", 200, payload)
}

func (o *GetAssetsDocumentationMrnOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/documentation/{mrn}][%d] getAssetsDocumentationMrnOK %s", 200, payload)
}

func (o *GetAssetsDocumentationMrnOK) GetPayload() []*models.AssetdocsDocumentation {
	return o.Payload
}

func (o *GetAssetsDocumentationMrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetsDocumentationMrnNotFound creates a GetAssetsDocumentationMrnNotFound with default headers values
func NewGetAssetsDocumentationMrnNotFound() *GetAssetsDocumentationMrnNotFound {
	return &GetAssetsDocumentationMrnNotFound{}
}

/*
GetAssetsDocumentationMrnNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAssetsDocumentationMrnNotFound struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get assets documentation mrn not found response has a 2xx status code
func (o *GetAssetsDocumentationMrnNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assets documentation mrn not found response has a 3xx status code
func (o *GetAssetsDocumentationMrnNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets documentation mrn not found response has a 4xx status code
func (o *GetAssetsDocumentationMrnNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assets documentation mrn not found response has a 5xx status code
func (o *GetAssetsDocumentationMrnNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get assets documentation mrn not found response a status code equal to that given
func (o *GetAssetsDocumentationMrnNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get assets documentation mrn not found response
func (o *GetAssetsDocumentationMrnNotFound) Code() int {
	return 404
}

func (o *GetAssetsDocumentationMrnNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/documentation/{mrn}][%d] getAssetsDocumentationMrnNotFound %s", 404, payload)
}

func (o *GetAssetsDocumentationMrnNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/documentation/{mrn}][%d] getAssetsDocumentationMrnNotFound %s", 404, payload)
}

func (o *GetAssetsDocumentationMrnNotFound) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetAssetsDocumentationMrnNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetsDocumentationMrnInternalServerError creates a GetAssetsDocumentationMrnInternalServerError with default headers values
func NewGetAssetsDocumentationMrnInternalServerError() *GetAssetsDocumentationMrnInternalServerError {
	return &GetAssetsDocumentationMrnInternalServerError{}
}

/*
GetAssetsDocumentationMrnInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAssetsDocumentationMrnInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get assets documentation mrn internal server error response has a 2xx status code
func (o *GetAssetsDocumentationMrnInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assets documentation mrn internal server error response has a 3xx status code
func (o *GetAssetsDocumentationMrnInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assets documentation mrn internal server error response has a 4xx status code
func (o *GetAssetsDocumentationMrnInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assets documentation mrn internal server error response has a 5xx status code
func (o *GetAssetsDocumentationMrnInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get assets documentation mrn internal server error response a status code equal to that given
func (o *GetAssetsDocumentationMrnInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get assets documentation mrn internal server error response
func (o *GetAssetsDocumentationMrnInternalServerError) Code() int {
	return 500
}

func (o *GetAssetsDocumentationMrnInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/documentation/{mrn}][%d] getAssetsDocumentationMrnInternalServerError %s", 500, payload)
}

func (o *GetAssetsDocumentationMrnInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /assets/documentation/{mrn}][%d] getAssetsDocumentationMrnInternalServerError %s", 500, payload)
}

func (o *GetAssetsDocumentationMrnInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetAssetsDocumentationMrnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
