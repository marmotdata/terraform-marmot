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

// PostAssetsBatchReader is a Reader for the PostAssetsBatch structure.
type PostAssetsBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAssetsBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAssetsBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAssetsBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAssetsBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAssetsBatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /assets/batch] PostAssetsBatch", response, response.Code())
	}
}

// NewPostAssetsBatchOK creates a PostAssetsBatchOK with default headers values
func NewPostAssetsBatchOK() *PostAssetsBatchOK {
	return &PostAssetsBatchOK{}
}

/*
PostAssetsBatchOK describes a response with status code 200, with default header values.

OK
*/
type PostAssetsBatchOK struct {
	Payload *models.AssetsBatchCreateResponse
}

// IsSuccess returns true when this post assets batch o k response has a 2xx status code
func (o *PostAssetsBatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post assets batch o k response has a 3xx status code
func (o *PostAssetsBatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post assets batch o k response has a 4xx status code
func (o *PostAssetsBatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post assets batch o k response has a 5xx status code
func (o *PostAssetsBatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post assets batch o k response a status code equal to that given
func (o *PostAssetsBatchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post assets batch o k response
func (o *PostAssetsBatchOK) Code() int {
	return 200
}

func (o *PostAssetsBatchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchOK %s", 200, payload)
}

func (o *PostAssetsBatchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchOK %s", 200, payload)
}

func (o *PostAssetsBatchOK) GetPayload() *models.AssetsBatchCreateResponse {
	return o.Payload
}

func (o *PostAssetsBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetsBatchCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAssetsBatchBadRequest creates a PostAssetsBatchBadRequest with default headers values
func NewPostAssetsBatchBadRequest() *PostAssetsBatchBadRequest {
	return &PostAssetsBatchBadRequest{}
}

/*
PostAssetsBatchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAssetsBatchBadRequest struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this post assets batch bad request response has a 2xx status code
func (o *PostAssetsBatchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post assets batch bad request response has a 3xx status code
func (o *PostAssetsBatchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post assets batch bad request response has a 4xx status code
func (o *PostAssetsBatchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post assets batch bad request response has a 5xx status code
func (o *PostAssetsBatchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post assets batch bad request response a status code equal to that given
func (o *PostAssetsBatchBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post assets batch bad request response
func (o *PostAssetsBatchBadRequest) Code() int {
	return 400
}

func (o *PostAssetsBatchBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchBadRequest %s", 400, payload)
}

func (o *PostAssetsBatchBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchBadRequest %s", 400, payload)
}

func (o *PostAssetsBatchBadRequest) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *PostAssetsBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAssetsBatchUnauthorized creates a PostAssetsBatchUnauthorized with default headers values
func NewPostAssetsBatchUnauthorized() *PostAssetsBatchUnauthorized {
	return &PostAssetsBatchUnauthorized{}
}

/*
PostAssetsBatchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostAssetsBatchUnauthorized struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this post assets batch unauthorized response has a 2xx status code
func (o *PostAssetsBatchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post assets batch unauthorized response has a 3xx status code
func (o *PostAssetsBatchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post assets batch unauthorized response has a 4xx status code
func (o *PostAssetsBatchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post assets batch unauthorized response has a 5xx status code
func (o *PostAssetsBatchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post assets batch unauthorized response a status code equal to that given
func (o *PostAssetsBatchUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post assets batch unauthorized response
func (o *PostAssetsBatchUnauthorized) Code() int {
	return 401
}

func (o *PostAssetsBatchUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchUnauthorized %s", 401, payload)
}

func (o *PostAssetsBatchUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchUnauthorized %s", 401, payload)
}

func (o *PostAssetsBatchUnauthorized) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *PostAssetsBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAssetsBatchInternalServerError creates a PostAssetsBatchInternalServerError with default headers values
func NewPostAssetsBatchInternalServerError() *PostAssetsBatchInternalServerError {
	return &PostAssetsBatchInternalServerError{}
}

/*
PostAssetsBatchInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAssetsBatchInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this post assets batch internal server error response has a 2xx status code
func (o *PostAssetsBatchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post assets batch internal server error response has a 3xx status code
func (o *PostAssetsBatchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post assets batch internal server error response has a 4xx status code
func (o *PostAssetsBatchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post assets batch internal server error response has a 5xx status code
func (o *PostAssetsBatchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post assets batch internal server error response a status code equal to that given
func (o *PostAssetsBatchInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post assets batch internal server error response
func (o *PostAssetsBatchInternalServerError) Code() int {
	return 500
}

func (o *PostAssetsBatchInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchInternalServerError %s", 500, payload)
}

func (o *PostAssetsBatchInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /assets/batch][%d] postAssetsBatchInternalServerError %s", 500, payload)
}

func (o *PostAssetsBatchInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *PostAssetsBatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
