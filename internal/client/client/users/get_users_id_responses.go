// Code generated by go-swagger; DO NOT EDIT.

package users

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

// GetUsersIDReader is a Reader for the GetUsersID structure.
type GetUsersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUsersIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{id}] GetUsersID", response, response.Code())
	}
}

// NewGetUsersIDOK creates a GetUsersIDOK with default headers values
func NewGetUsersIDOK() *GetUsersIDOK {
	return &GetUsersIDOK{}
}

/*
GetUsersIDOK describes a response with status code 200, with default header values.

OK
*/
type GetUsersIDOK struct {
	Payload *models.UserUser
}

// IsSuccess returns true when this get users Id o k response has a 2xx status code
func (o *GetUsersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users Id o k response has a 3xx status code
func (o *GetUsersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users Id o k response has a 4xx status code
func (o *GetUsersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users Id o k response has a 5xx status code
func (o *GetUsersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users Id o k response a status code equal to that given
func (o *GetUsersIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users Id o k response
func (o *GetUsersIDOK) Code() int {
	return 200
}

func (o *GetUsersIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdOK %s", 200, payload)
}

func (o *GetUsersIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdOK %s", 200, payload)
}

func (o *GetUsersIDOK) GetPayload() *models.UserUser {
	return o.Payload
}

func (o *GetUsersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDNotFound creates a GetUsersIDNotFound with default headers values
func NewGetUsersIDNotFound() *GetUsersIDNotFound {
	return &GetUsersIDNotFound{}
}

/*
GetUsersIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUsersIDNotFound struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get users Id not found response has a 2xx status code
func (o *GetUsersIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users Id not found response has a 3xx status code
func (o *GetUsersIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users Id not found response has a 4xx status code
func (o *GetUsersIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users Id not found response has a 5xx status code
func (o *GetUsersIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get users Id not found response a status code equal to that given
func (o *GetUsersIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get users Id not found response
func (o *GetUsersIDNotFound) Code() int {
	return 404
}

func (o *GetUsersIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdNotFound %s", 404, payload)
}

func (o *GetUsersIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdNotFound %s", 404, payload)
}

func (o *GetUsersIDNotFound) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetUsersIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersIDInternalServerError creates a GetUsersIDInternalServerError with default headers values
func NewGetUsersIDInternalServerError() *GetUsersIDInternalServerError {
	return &GetUsersIDInternalServerError{}
}

/*
GetUsersIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUsersIDInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get users Id internal server error response has a 2xx status code
func (o *GetUsersIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users Id internal server error response has a 3xx status code
func (o *GetUsersIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users Id internal server error response has a 4xx status code
func (o *GetUsersIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users Id internal server error response has a 5xx status code
func (o *GetUsersIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get users Id internal server error response a status code equal to that given
func (o *GetUsersIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get users Id internal server error response
func (o *GetUsersIDInternalServerError) Code() int {
	return 500
}

func (o *GetUsersIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdInternalServerError %s", 500, payload)
}

func (o *GetUsersIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{id}][%d] getUsersIdInternalServerError %s", 500, payload)
}

func (o *GetUsersIDInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetUsersIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
