// Code generated by go-swagger; DO NOT EDIT.

package actions

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

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateActionsPropertiesUpdateParams creates a new WeaviateActionsPropertiesUpdateParams object
// with the default values initialized.
func NewWeaviateActionsPropertiesUpdateParams() *WeaviateActionsPropertiesUpdateParams {
	var ()
	return &WeaviateActionsPropertiesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsPropertiesUpdateParamsWithTimeout creates a new WeaviateActionsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsPropertiesUpdateParamsWithTimeout(timeout time.Duration) *WeaviateActionsPropertiesUpdateParams {
	var ()
	return &WeaviateActionsPropertiesUpdateParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsPropertiesUpdateParamsWithContext creates a new WeaviateActionsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsPropertiesUpdateParamsWithContext(ctx context.Context) *WeaviateActionsPropertiesUpdateParams {
	var ()
	return &WeaviateActionsPropertiesUpdateParams{

		Context: ctx,
	}
}

// NewWeaviateActionsPropertiesUpdateParamsWithHTTPClient creates a new WeaviateActionsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsPropertiesUpdateParamsWithHTTPClient(client *http.Client) *WeaviateActionsPropertiesUpdateParams {
	var ()
	return &WeaviateActionsPropertiesUpdateParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsPropertiesUpdateParams contains all the parameters to send to the API endpoint
for the weaviate actions properties update operation typically these are written to a http.Request
*/
type WeaviateActionsPropertiesUpdateParams struct {

	/*ActionID
	  Unique ID of the action.

	*/
	ActionID strfmt.UUID
	/*Body*/
	Body models.MultipleRef
	/*PropertyName
	  Unique name of the property related to the action.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) WithTimeout(timeout time.Duration) *WeaviateActionsPropertiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) WithContext(ctx context.Context) *WeaviateActionsPropertiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) WithHTTPClient(client *http.Client) *WeaviateActionsPropertiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) WithActionID(actionID strfmt.UUID) *WeaviateActionsPropertiesUpdateParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WithBody adds the body to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) WithBody(body models.MultipleRef) *WeaviateActionsPropertiesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) SetBody(body models.MultipleRef) {
	o.Body = body
}

// WithPropertyName adds the propertyName to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) WithPropertyName(propertyName string) *WeaviateActionsPropertiesUpdateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate actions properties update params
func (o *WeaviateActionsPropertiesUpdateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsPropertiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}