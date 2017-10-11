// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSMTPTemplateParams creates a new GetSMTPTemplateParams object
// with the default values initialized.
func NewGetSMTPTemplateParams() *GetSMTPTemplateParams {
	var ()
	return &GetSMTPTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSMTPTemplateParamsWithTimeout creates a new GetSMTPTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSMTPTemplateParamsWithTimeout(timeout time.Duration) *GetSMTPTemplateParams {
	var ()
	return &GetSMTPTemplateParams{

		timeout: timeout,
	}
}

// NewGetSMTPTemplateParamsWithContext creates a new GetSMTPTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSMTPTemplateParamsWithContext(ctx context.Context) *GetSMTPTemplateParams {
	var ()
	return &GetSMTPTemplateParams{

		Context: ctx,
	}
}

// NewGetSMTPTemplateParamsWithHTTPClient creates a new GetSMTPTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSMTPTemplateParamsWithHTTPClient(client *http.Client) *GetSMTPTemplateParams {
	var ()
	return &GetSMTPTemplateParams{
		HTTPClient: client,
	}
}

/*GetSMTPTemplateParams contains all the parameters to send to the API endpoint
for the get Smtp template operation typically these are written to a http.Request
*/
type GetSMTPTemplateParams struct {

	/*TemplateID
	  id of the template

	*/
	TemplateID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Smtp template params
func (o *GetSMTPTemplateParams) WithTimeout(timeout time.Duration) *GetSMTPTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Smtp template params
func (o *GetSMTPTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Smtp template params
func (o *GetSMTPTemplateParams) WithContext(ctx context.Context) *GetSMTPTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Smtp template params
func (o *GetSMTPTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Smtp template params
func (o *GetSMTPTemplateParams) WithHTTPClient(client *http.Client) *GetSMTPTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Smtp template params
func (o *GetSMTPTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplateID adds the templateID to the get Smtp template params
func (o *GetSMTPTemplateParams) WithTemplateID(templateID int64) *GetSMTPTemplateParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the get Smtp template params
func (o *GetSMTPTemplateParams) SetTemplateID(templateID int64) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSMTPTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param templateId
	if err := r.SetPathParam("templateId", swag.FormatInt64(o.TemplateID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}