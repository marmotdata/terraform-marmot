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

// NewPostAssetsParams creates a new PostAssetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAssetsParams() *PostAssetsParams {
	return &PostAssetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAssetsParamsWithTimeout creates a new PostAssetsParams object
// with the ability to set a timeout on a request.
func NewPostAssetsParamsWithTimeout(timeout time.Duration) *PostAssetsParams {
	return &PostAssetsParams{
		timeout: timeout,
	}
}

// NewPostAssetsParamsWithContext creates a new PostAssetsParams object
// with the ability to set a context for a request.
func NewPostAssetsParamsWithContext(ctx context.Context) *PostAssetsParams {
	return &PostAssetsParams{
		Context: ctx,
	}
}

// NewPostAssetsParamsWithHTTPClient creates a new PostAssetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAssetsParamsWithHTTPClient(client *http.Client) *PostAssetsParams {
	return &PostAssetsParams{
		HTTPClient: client,
	}
}

/*
PostAssetsParams contains all the parameters to send to the API endpoint

	for the post assets operation.

	Typically these are written to a http.Request.
*/
type PostAssetsParams struct {

	/* Asset.

	   Asset creation request
	*/
	Asset *models.AssetsCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post assets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAssetsParams) WithDefaults() *PostAssetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post assets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAssetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post assets params
func (o *PostAssetsParams) WithTimeout(timeout time.Duration) *PostAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post assets params
func (o *PostAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post assets params
func (o *PostAssetsParams) WithContext(ctx context.Context) *PostAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post assets params
func (o *PostAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post assets params
func (o *PostAssetsParams) WithHTTPClient(client *http.Client) *PostAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post assets params
func (o *PostAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsset adds the asset to the post assets params
func (o *PostAssetsParams) WithAsset(asset *models.AssetsCreateRequest) *PostAssetsParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the post assets params
func (o *PostAssetsParams) SetAsset(asset *models.AssetsCreateRequest) {
	o.Asset = asset
}

// WriteToRequest writes these params to a swagger request
func (o *PostAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Asset != nil {
		if err := r.SetBodyParam(o.Asset); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
