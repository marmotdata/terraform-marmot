// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/marmotdata/terraform-provider-marmot/internal/client/models"
)

// NewDeleteAssetsIDTagsParams creates a new DeleteAssetsIDTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAssetsIDTagsParams() *DeleteAssetsIDTagsParams {
	return &DeleteAssetsIDTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAssetsIDTagsParamsWithTimeout creates a new DeleteAssetsIDTagsParams object
// with the ability to set a timeout on a request.
func NewDeleteAssetsIDTagsParamsWithTimeout(timeout time.Duration) *DeleteAssetsIDTagsParams {
	return &DeleteAssetsIDTagsParams{
		timeout: timeout,
	}
}

// NewDeleteAssetsIDTagsParamsWithContext creates a new DeleteAssetsIDTagsParams object
// with the ability to set a context for a request.
func NewDeleteAssetsIDTagsParamsWithContext(ctx context.Context) *DeleteAssetsIDTagsParams {
	return &DeleteAssetsIDTagsParams{
		Context: ctx,
	}
}

// NewDeleteAssetsIDTagsParamsWithHTTPClient creates a new DeleteAssetsIDTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAssetsIDTagsParamsWithHTTPClient(client *http.Client) *DeleteAssetsIDTagsParams {
	return &DeleteAssetsIDTagsParams{
		HTTPClient: client,
	}
}

/*
DeleteAssetsIDTagsParams contains all the parameters to send to the API endpoint

	for the delete assets ID tags operation.

	Typically these are written to a http.Request.
*/
type DeleteAssetsIDTagsParams struct {

	/* ID.

	   Asset ID
	*/
	ID string

	/* Tag.

	   Tag to remove
	*/
	Tag *models.AssetsTagRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete assets ID tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAssetsIDTagsParams) WithDefaults() *DeleteAssetsIDTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete assets ID tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAssetsIDTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) WithTimeout(timeout time.Duration) *DeleteAssetsIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) WithContext(ctx context.Context) *DeleteAssetsIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) WithHTTPClient(client *http.Client) *DeleteAssetsIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) WithID(id string) *DeleteAssetsIDTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) SetID(id string) {
	o.ID = id
}

// WithTag adds the tag to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) WithTag(tag *models.AssetsTagRequest) *DeleteAssetsIDTagsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the delete assets ID tags params
func (o *DeleteAssetsIDTagsParams) SetTag(tag *models.AssetsTagRequest) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAssetsIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
