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

// DeleteUsersApikeysIDReader is a Reader for the DeleteUsersApikeysID structure.
type DeleteUsersApikeysIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersApikeysIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUsersApikeysIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUsersApikeysIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /users/apikeys/{id}] DeleteUsersApikeysID", response, response.Code())
	}
}

// NewDeleteUsersApikeysIDNoContent creates a DeleteUsersApikeysIDNoContent with default headers values
func NewDeleteUsersApikeysIDNoContent() *DeleteUsersApikeysIDNoContent {
	return &DeleteUsersApikeysIDNoContent{}
}

/*
DeleteUsersApikeysIDNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUsersApikeysIDNoContent struct {
}

// IsSuccess returns true when this delete users apikeys Id no content response has a 2xx status code
func (o *DeleteUsersApikeysIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete users apikeys Id no content response has a 3xx status code
func (o *DeleteUsersApikeysIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete users apikeys Id no content response has a 4xx status code
func (o *DeleteUsersApikeysIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete users apikeys Id no content response has a 5xx status code
func (o *DeleteUsersApikeysIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete users apikeys Id no content response a status code equal to that given
func (o *DeleteUsersApikeysIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete users apikeys Id no content response
func (o *DeleteUsersApikeysIDNoContent) Code() int {
	return 204
}

func (o *DeleteUsersApikeysIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/apikeys/{id}][%d] deleteUsersApikeysIdNoContent", 204)
}

func (o *DeleteUsersApikeysIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/apikeys/{id}][%d] deleteUsersApikeysIdNoContent", 204)
}

func (o *DeleteUsersApikeysIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersApikeysIDBadRequest creates a DeleteUsersApikeysIDBadRequest with default headers values
func NewDeleteUsersApikeysIDBadRequest() *DeleteUsersApikeysIDBadRequest {
	return &DeleteUsersApikeysIDBadRequest{}
}

/*
DeleteUsersApikeysIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUsersApikeysIDBadRequest struct {
	Payload *models.CommonErrorResponse
}

// IsSuccess returns true when this delete users apikeys Id bad request response has a 2xx status code
func (o *DeleteUsersApikeysIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete users apikeys Id bad request response has a 3xx status code
func (o *DeleteUsersApikeysIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete users apikeys Id bad request response has a 4xx status code
func (o *DeleteUsersApikeysIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete users apikeys Id bad request response has a 5xx status code
func (o *DeleteUsersApikeysIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete users apikeys Id bad request response a status code equal to that given
func (o *DeleteUsersApikeysIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete users apikeys Id bad request response
func (o *DeleteUsersApikeysIDBadRequest) Code() int {
	return 400
}

func (o *DeleteUsersApikeysIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/apikeys/{id}][%d] deleteUsersApikeysIdBadRequest %s", 400, payload)
}

func (o *DeleteUsersApikeysIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/apikeys/{id}][%d] deleteUsersApikeysIdBadRequest %s", 400, payload)
}

func (o *DeleteUsersApikeysIDBadRequest) GetPayload() *models.CommonErrorResponse {
	return o.Payload
}

func (o *DeleteUsersApikeysIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
