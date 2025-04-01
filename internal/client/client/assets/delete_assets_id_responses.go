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

// DeleteAssetsIDReader is a Reader for the DeleteAssetsID structure.
type DeleteAssetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAssetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAssetsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAssetsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAssetsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAssetsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /assets/{id}] DeleteAssetsID", response, response.Code())
	}
}

// NewDeleteAssetsIDNoContent creates a DeleteAssetsIDNoContent with default headers values
func NewDeleteAssetsIDNoContent() *DeleteAssetsIDNoContent {
	return &DeleteAssetsIDNoContent{}
}

/*
DeleteAssetsIDNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteAssetsIDNoContent struct {
}

// IsSuccess returns true when this delete assets Id no content response has a 2xx status code
func (o *DeleteAssetsIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete assets Id no content response has a 3xx status code
func (o *DeleteAssetsIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete assets Id no content response has a 4xx status code
func (o *DeleteAssetsIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete assets Id no content response has a 5xx status code
func (o *DeleteAssetsIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete assets Id no content response a status code equal to that given
func (o *DeleteAssetsIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete assets Id no content response
func (o *DeleteAssetsIDNoContent) Code() int {
	return 204
}

func (o *DeleteAssetsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdNoContent", 204)
}

func (o *DeleteAssetsIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdNoContent", 204)
}

func (o *DeleteAssetsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAssetsIDNotFound creates a DeleteAssetsIDNotFound with default headers values
func NewDeleteAssetsIDNotFound() *DeleteAssetsIDNotFound {
	return &DeleteAssetsIDNotFound{}
}

/*
DeleteAssetsIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAssetsIDNotFound struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this delete assets Id not found response has a 2xx status code
func (o *DeleteAssetsIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete assets Id not found response has a 3xx status code
func (o *DeleteAssetsIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete assets Id not found response has a 4xx status code
func (o *DeleteAssetsIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete assets Id not found response has a 5xx status code
func (o *DeleteAssetsIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete assets Id not found response a status code equal to that given
func (o *DeleteAssetsIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete assets Id not found response
func (o *DeleteAssetsIDNotFound) Code() int {
	return 404
}

func (o *DeleteAssetsIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdNotFound %s", 404, payload)
}

func (o *DeleteAssetsIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdNotFound %s", 404, payload)
}

func (o *DeleteAssetsIDNotFound) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *DeleteAssetsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAssetsIDConflict creates a DeleteAssetsIDConflict with default headers values
func NewDeleteAssetsIDConflict() *DeleteAssetsIDConflict {
	return &DeleteAssetsIDConflict{}
}

/*
DeleteAssetsIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAssetsIDConflict struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this delete assets Id conflict response has a 2xx status code
func (o *DeleteAssetsIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete assets Id conflict response has a 3xx status code
func (o *DeleteAssetsIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete assets Id conflict response has a 4xx status code
func (o *DeleteAssetsIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete assets Id conflict response has a 5xx status code
func (o *DeleteAssetsIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete assets Id conflict response a status code equal to that given
func (o *DeleteAssetsIDConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete assets Id conflict response
func (o *DeleteAssetsIDConflict) Code() int {
	return 409
}

func (o *DeleteAssetsIDConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdConflict %s", 409, payload)
}

func (o *DeleteAssetsIDConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdConflict %s", 409, payload)
}

func (o *DeleteAssetsIDConflict) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *DeleteAssetsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAssetsIDInternalServerError creates a DeleteAssetsIDInternalServerError with default headers values
func NewDeleteAssetsIDInternalServerError() *DeleteAssetsIDInternalServerError {
	return &DeleteAssetsIDInternalServerError{}
}

/*
DeleteAssetsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAssetsIDInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this delete assets Id internal server error response has a 2xx status code
func (o *DeleteAssetsIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete assets Id internal server error response has a 3xx status code
func (o *DeleteAssetsIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete assets Id internal server error response has a 4xx status code
func (o *DeleteAssetsIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete assets Id internal server error response has a 5xx status code
func (o *DeleteAssetsIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete assets Id internal server error response a status code equal to that given
func (o *DeleteAssetsIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete assets Id internal server error response
func (o *DeleteAssetsIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteAssetsIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdInternalServerError %s", 500, payload)
}

func (o *DeleteAssetsIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /assets/{id}][%d] deleteAssetsIdInternalServerError %s", 500, payload)
}

func (o *DeleteAssetsIDInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *DeleteAssetsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
