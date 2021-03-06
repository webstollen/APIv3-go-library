// Code generated by go-swagger; DO NOT EDIT.

package transactional_sms

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

// NewGetTransacSMSReportParams creates a new GetTransacSMSReportParams object
// with the default values initialized.
func NewGetTransacSMSReportParams() *GetTransacSMSReportParams {
	var ()
	return &GetTransacSMSReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransacSMSReportParamsWithTimeout creates a new GetTransacSMSReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTransacSMSReportParamsWithTimeout(timeout time.Duration) *GetTransacSMSReportParams {
	var ()
	return &GetTransacSMSReportParams{

		timeout: timeout,
	}
}

// NewGetTransacSMSReportParamsWithContext creates a new GetTransacSMSReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTransacSMSReportParamsWithContext(ctx context.Context) *GetTransacSMSReportParams {
	var ()
	return &GetTransacSMSReportParams{

		Context: ctx,
	}
}

// NewGetTransacSMSReportParamsWithHTTPClient creates a new GetTransacSMSReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTransacSMSReportParamsWithHTTPClient(client *http.Client) *GetTransacSMSReportParams {
	var ()
	return &GetTransacSMSReportParams{
		HTTPClient: client,
	}
}

/*GetTransacSMSReportParams contains all the parameters to send to the API endpoint
for the get transac Sms report operation typically these are written to a http.Request
*/
type GetTransacSMSReportParams struct {

	/*Days
	  Number of days in the past including today (positive integer). Not compatible with 'startDate' and 'endDate'

	*/
	Days *int64
	/*EndDate
	  Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report

	*/
	EndDate *strfmt.Date
	/*StartDate
	  Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report

	*/
	StartDate *strfmt.Date
	/*Tag
	  Filter on a tag

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithTimeout(timeout time.Duration) *GetTransacSMSReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithContext(ctx context.Context) *GetTransacSMSReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithHTTPClient(client *http.Client) *GetTransacSMSReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDays adds the days to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithDays(days *int64) *GetTransacSMSReportParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetDays(days *int64) {
	o.Days = days
}

// WithEndDate adds the endDate to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithEndDate(endDate *strfmt.Date) *GetTransacSMSReportParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithStartDate(startDate *strfmt.Date) *GetTransacSMSReportParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithTag adds the tag to the get transac Sms report params
func (o *GetTransacSMSReportParams) WithTag(tag *string) *GetTransacSMSReportParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get transac Sms report params
func (o *GetTransacSMSReportParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransacSMSReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Days != nil {

		// query param days
		var qrDays int64
		if o.Days != nil {
			qrDays = *o.Days
		}
		qDays := swag.FormatInt64(qrDays)
		if qDays != "" {
			if err := r.SetQueryParam("days", qDays); err != nil {
				return err
			}
		}

	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
