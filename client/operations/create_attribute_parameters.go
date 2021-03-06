// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewCreateAttributeParams creates a new CreateAttributeParams object
// with the default values initialized.
func NewCreateAttributeParams() *CreateAttributeParams {
	var ()
	return &CreateAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAttributeParamsWithTimeout creates a new CreateAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAttributeParamsWithTimeout(timeout time.Duration) *CreateAttributeParams {
	var ()
	return &CreateAttributeParams{

		timeout: timeout,
	}
}

// NewCreateAttributeParamsWithContext creates a new CreateAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAttributeParamsWithContext(ctx context.Context) *CreateAttributeParams {
	var ()
	return &CreateAttributeParams{

		Context: ctx,
	}
}

// NewCreateAttributeParamsWithHTTPClient creates a new CreateAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAttributeParamsWithHTTPClient(client *http.Client) *CreateAttributeParams {
	var ()
	return &CreateAttributeParams{
		HTTPClient: client,
	}
}

/*CreateAttributeParams contains all the parameters to send to the API endpoint
for the create attribute operation typically these are written to a http.Request
*/
type CreateAttributeParams struct {

	/*AttributeCategory
	  Category of the attribute

	*/
	AttributeCategory string
	/*AttributeName
	  Name of the attribute

	*/
	AttributeName string
	/*CreateAttribute
	  Values to create an attribute

	*/
	CreateAttribute *models.CreateAttribute

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create attribute params
func (o *CreateAttributeParams) WithTimeout(timeout time.Duration) *CreateAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create attribute params
func (o *CreateAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create attribute params
func (o *CreateAttributeParams) WithContext(ctx context.Context) *CreateAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create attribute params
func (o *CreateAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create attribute params
func (o *CreateAttributeParams) WithHTTPClient(client *http.Client) *CreateAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create attribute params
func (o *CreateAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeCategory adds the attributeCategory to the create attribute params
func (o *CreateAttributeParams) WithAttributeCategory(attributeCategory string) *CreateAttributeParams {
	o.SetAttributeCategory(attributeCategory)
	return o
}

// SetAttributeCategory adds the attributeCategory to the create attribute params
func (o *CreateAttributeParams) SetAttributeCategory(attributeCategory string) {
	o.AttributeCategory = attributeCategory
}

// WithAttributeName adds the attributeName to the create attribute params
func (o *CreateAttributeParams) WithAttributeName(attributeName string) *CreateAttributeParams {
	o.SetAttributeName(attributeName)
	return o
}

// SetAttributeName adds the attributeName to the create attribute params
func (o *CreateAttributeParams) SetAttributeName(attributeName string) {
	o.AttributeName = attributeName
}

// WithCreateAttribute adds the createAttribute to the create attribute params
func (o *CreateAttributeParams) WithCreateAttribute(createAttribute *models.CreateAttribute) *CreateAttributeParams {
	o.SetCreateAttribute(createAttribute)
	return o
}

// SetCreateAttribute adds the createAttribute to the create attribute params
func (o *CreateAttributeParams) SetCreateAttribute(createAttribute *models.CreateAttribute) {
	o.CreateAttribute = createAttribute
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeCategory
	if err := r.SetPathParam("attributeCategory", o.AttributeCategory); err != nil {
		return err
	}

	// path param attributeName
	if err := r.SetPathParam("attributeName", o.AttributeName); err != nil {
		return err
	}

	if o.CreateAttribute != nil {
		if err := r.SetBodyParam(o.CreateAttribute); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
