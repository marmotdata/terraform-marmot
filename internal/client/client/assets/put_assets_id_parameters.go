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

// NewPutAssetsIDParams creates a new PutAssetsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAssetsIDParams() *PutAssetsIDParams {
	return &PutAssetsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAssetsIDParamsWithTimeout creates a new PutAssetsIDParams object
// with the ability to set a timeout on a request.
func NewPutAssetsIDParamsWithTimeout(timeout time.Duration) *PutAssetsIDParams {
	return &PutAssetsIDParams{
		timeout: timeout,
	}
}

// NewPutAssetsIDParamsWithContext creates a new PutAssetsIDParams object
// with the ability to set a context for a request.
func NewPutAssetsIDParamsWithContext(ctx context.Context) *PutAssetsIDParams {
	return &PutAssetsIDParams{
		Context: ctx,
	}
}

// NewPutAssetsIDParamsWithHTTPClient creates a new PutAssetsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAssetsIDParamsWithHTTPClient(client *http.Client) *PutAssetsIDParams {
	return &PutAssetsIDParams{
		HTTPClient: client,
	}
}

/*
PutAssetsIDParams contains all the parameters to send to the API endpoint

	for the put assets ID operation.

	Typically these are written to a http.Request.
*/
type PutAssetsIDParams struct {

	/* Asset.

	   Asset update request
	*/
	Asset *models.AssetsUpdateRequest

	/* ID.

	   Asset ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put assets ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAssetsIDParams) WithDefaults() *PutAssetsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put assets ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAssetsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put assets ID params
func (o *PutAssetsIDParams) WithTimeout(timeout time.Duration) *PutAssetsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put assets ID params
func (o *PutAssetsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put assets ID params
func (o *PutAssetsIDParams) WithContext(ctx context.Context) *PutAssetsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put assets ID params
func (o *PutAssetsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put assets ID params
func (o *PutAssetsIDParams) WithHTTPClient(client *http.Client) *PutAssetsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put assets ID params
func (o *PutAssetsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsset adds the asset to the put assets ID params
func (o *PutAssetsIDParams) WithAsset(asset *models.AssetsUpdateRequest) *PutAssetsIDParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the put assets ID params
func (o *PutAssetsIDParams) SetAsset(asset *models.AssetsUpdateRequest) {
	o.Asset = asset
}

// WithID adds the id to the put assets ID params
func (o *PutAssetsIDParams) WithID(id string) *PutAssetsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put assets ID params
func (o *PutAssetsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutAssetsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Asset != nil {
		if err := r.SetBodyParam(o.Asset); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
