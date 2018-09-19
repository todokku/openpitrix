// Code generated by go-swagger; DO NOT EDIT.

package account_manager

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

// NewDescribeUsersParams creates a new DescribeUsersParams object
// with the default values initialized.
func NewDescribeUsersParams() *DescribeUsersParams {
	var ()
	return &DescribeUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeUsersParamsWithTimeout creates a new DescribeUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeUsersParamsWithTimeout(timeout time.Duration) *DescribeUsersParams {
	var ()
	return &DescribeUsersParams{

		timeout: timeout,
	}
}

// NewDescribeUsersParamsWithContext creates a new DescribeUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeUsersParamsWithContext(ctx context.Context) *DescribeUsersParams {
	var ()
	return &DescribeUsersParams{

		Context: ctx,
	}
}

// NewDescribeUsersParamsWithHTTPClient creates a new DescribeUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeUsersParamsWithHTTPClient(client *http.Client) *DescribeUsersParams {
	var ()
	return &DescribeUsersParams{
		HTTPClient: client,
	}
}

/*DescribeUsersParams contains all the parameters to send to the API endpoint
for the describe users operation typically these are written to a http.Request
*/
type DescribeUsersParams struct {

	/*GroupID*/
	GroupID []string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Reverse*/
	Reverse *bool
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string
	/*UserID*/
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe users params
func (o *DescribeUsersParams) WithTimeout(timeout time.Duration) *DescribeUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe users params
func (o *DescribeUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe users params
func (o *DescribeUsersParams) WithContext(ctx context.Context) *DescribeUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe users params
func (o *DescribeUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe users params
func (o *DescribeUsersParams) WithHTTPClient(client *http.Client) *DescribeUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe users params
func (o *DescribeUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the describe users params
func (o *DescribeUsersParams) WithGroupID(groupID []string) *DescribeUsersParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the describe users params
func (o *DescribeUsersParams) SetGroupID(groupID []string) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the describe users params
func (o *DescribeUsersParams) WithLimit(limit *int64) *DescribeUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe users params
func (o *DescribeUsersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe users params
func (o *DescribeUsersParams) WithOffset(offset *int64) *DescribeUsersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe users params
func (o *DescribeUsersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithReverse adds the reverse to the describe users params
func (o *DescribeUsersParams) WithReverse(reverse *bool) *DescribeUsersParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe users params
func (o *DescribeUsersParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe users params
func (o *DescribeUsersParams) WithSearchWord(searchWord *string) *DescribeUsersParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe users params
func (o *DescribeUsersParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe users params
func (o *DescribeUsersParams) WithSortKey(sortKey *string) *DescribeUsersParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe users params
func (o *DescribeUsersParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe users params
func (o *DescribeUsersParams) WithStatus(status []string) *DescribeUsersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe users params
func (o *DescribeUsersParams) SetStatus(status []string) {
	o.Status = status
}

// WithUserID adds the userID to the describe users params
func (o *DescribeUsersParams) WithUserID(userID []string) *DescribeUsersParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the describe users params
func (o *DescribeUsersParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesGroupID := o.GroupID

	joinedGroupID := swag.JoinByFormat(valuesGroupID, "multi")
	// query array param group_id
	if err := r.SetQueryParam("group_id", joinedGroupID...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesUserID := o.UserID

	joinedUserID := swag.JoinByFormat(valuesUserID, "multi")
	// query array param user_id
	if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
