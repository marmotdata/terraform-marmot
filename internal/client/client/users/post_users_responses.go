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

// PostUsersReader is a Reader for the PostUsers structure.
type PostUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostUsersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /users] PostUsers", response, response.Code())
	}
}

// NewPostUsersOK creates a PostUsersOK with default headers values
func NewPostUsersOK() *PostUsersOK {
	return &PostUsersOK{}
}

/*
PostUsersOK describes a response with status code 200, with default header values.

OK
*/
type PostUsersOK struct {
	Payload *models.UserUser
}

// IsSuccess returns true when this post users o k response has a 2xx status code
func (o *PostUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post users o k response has a 3xx status code
func (o *PostUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users o k response has a 4xx status code
func (o *PostUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users o k response has a 5xx status code
func (o *PostUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post users o k response a status code equal to that given
func (o *PostUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post users o k response
func (o *PostUsersOK) Code() int {
	return 200
}

func (o *PostUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersOK %s", 200, payload)
}

func (o *PostUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersOK %s", 200, payload)
}

func (o *PostUsersOK) GetPayload() *models.UserUser {
	return o.Payload
}

func (o *PostUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersBadRequest creates a PostUsersBadRequest with default headers values
func NewPostUsersBadRequest() *PostUsersBadRequest {
	return &PostUsersBadRequest{}
}

/*
PostUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostUsersBadRequest struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this post users bad request response has a 2xx status code
func (o *PostUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users bad request response has a 3xx status code
func (o *PostUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users bad request response has a 4xx status code
func (o *PostUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users bad request response has a 5xx status code
func (o *PostUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post users bad request response a status code equal to that given
func (o *PostUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post users bad request response
func (o *PostUsersBadRequest) Code() int {
	return 400
}

func (o *PostUsersBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersBadRequest %s", 400, payload)
}

func (o *PostUsersBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersBadRequest %s", 400, payload)
}

func (o *PostUsersBadRequest) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *PostUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersConflict creates a PostUsersConflict with default headers values
func NewPostUsersConflict() *PostUsersConflict {
	return &PostUsersConflict{}
}

/*
PostUsersConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostUsersConflict struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this post users conflict response has a 2xx status code
func (o *PostUsersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users conflict response has a 3xx status code
func (o *PostUsersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users conflict response has a 4xx status code
func (o *PostUsersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users conflict response has a 5xx status code
func (o *PostUsersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post users conflict response a status code equal to that given
func (o *PostUsersConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post users conflict response
func (o *PostUsersConflict) Code() int {
	return 409
}

func (o *PostUsersConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersConflict %s", 409, payload)
}

func (o *PostUsersConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersConflict %s", 409, payload)
}

func (o *PostUsersConflict) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *PostUsersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
