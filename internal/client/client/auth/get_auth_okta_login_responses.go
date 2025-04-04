// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// GetAuthOktaLoginReader is a Reader for the GetAuthOktaLogin structure.
type GetAuthOktaLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthOktaLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 307:
		result := NewGetAuthOktaLoginTemporaryRedirect()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthOktaLoginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /auth/okta/login] GetAuthOktaLogin", response, response.Code())
	}
}

// NewGetAuthOktaLoginTemporaryRedirect creates a GetAuthOktaLoginTemporaryRedirect with default headers values
func NewGetAuthOktaLoginTemporaryRedirect() *GetAuthOktaLoginTemporaryRedirect {
	return &GetAuthOktaLoginTemporaryRedirect{}
}

/*
GetAuthOktaLoginTemporaryRedirect describes a response with status code 307, with default header values.

Temporary Redirect
*/
type GetAuthOktaLoginTemporaryRedirect struct {
	Payload string
}

// IsSuccess returns true when this get auth okta login temporary redirect response has a 2xx status code
func (o *GetAuthOktaLoginTemporaryRedirect) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth okta login temporary redirect response has a 3xx status code
func (o *GetAuthOktaLoginTemporaryRedirect) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get auth okta login temporary redirect response has a 4xx status code
func (o *GetAuthOktaLoginTemporaryRedirect) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth okta login temporary redirect response has a 5xx status code
func (o *GetAuthOktaLoginTemporaryRedirect) IsServerError() bool {
	return false
}

// IsCode returns true when this get auth okta login temporary redirect response a status code equal to that given
func (o *GetAuthOktaLoginTemporaryRedirect) IsCode(code int) bool {
	return code == 307
}

// Code gets the status code for the get auth okta login temporary redirect response
func (o *GetAuthOktaLoginTemporaryRedirect) Code() int {
	return 307
}

func (o *GetAuthOktaLoginTemporaryRedirect) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/okta/login][%d] getAuthOktaLoginTemporaryRedirect %s", 307, payload)
}

func (o *GetAuthOktaLoginTemporaryRedirect) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/okta/login][%d] getAuthOktaLoginTemporaryRedirect %s", 307, payload)
}

func (o *GetAuthOktaLoginTemporaryRedirect) GetPayload() string {
	return o.Payload
}

func (o *GetAuthOktaLoginTemporaryRedirect) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthOktaLoginInternalServerError creates a GetAuthOktaLoginInternalServerError with default headers values
func NewGetAuthOktaLoginInternalServerError() *GetAuthOktaLoginInternalServerError {
	return &GetAuthOktaLoginInternalServerError{}
}

/*
GetAuthOktaLoginInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAuthOktaLoginInternalServerError struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this get auth okta login internal server error response has a 2xx status code
func (o *GetAuthOktaLoginInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth okta login internal server error response has a 3xx status code
func (o *GetAuthOktaLoginInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get auth okta login internal server error response has a 4xx status code
func (o *GetAuthOktaLoginInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth okta login internal server error response has a 5xx status code
func (o *GetAuthOktaLoginInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get auth okta login internal server error response a status code equal to that given
func (o *GetAuthOktaLoginInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get auth okta login internal server error response
func (o *GetAuthOktaLoginInternalServerError) Code() int {
	return 500
}

func (o *GetAuthOktaLoginInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/okta/login][%d] getAuthOktaLoginInternalServerError %s", 500, payload)
}

func (o *GetAuthOktaLoginInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/okta/login][%d] getAuthOktaLoginInternalServerError %s", 500, payload)
}

func (o *GetAuthOktaLoginInternalServerError) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *GetAuthOktaLoginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
