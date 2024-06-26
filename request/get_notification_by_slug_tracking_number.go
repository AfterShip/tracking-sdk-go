// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
    "errors"
	"fmt"
	"net/http"
	"github.com/google/go-querystring/query"
	"github.com/aftership/tracking-sdk-go/v4/component"
	"github.com/aftership/tracking-sdk-go/v4/errorx"
	"github.com/aftership/tracking-sdk-go/v4/model"
)

type GetNotificationBySlugTrackingNumberRequest struct {
	slug     string
	trackingNumber     string
    query  model.GetNotificationBySlugTrackingNumberQuery
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewGetNotificationBySlugTrackingNumberRequest(sender *component.HttpSender) *GetNotificationBySlugTrackingNumberRequest {
	return &GetNotificationBySlugTrackingNumberRequest{
        sender: sender,
	}
}

func (t *GetNotificationBySlugTrackingNumberRequest) BuildQuery(query model.GetNotificationBySlugTrackingNumberQuery) *GetNotificationBySlugTrackingNumberRequest {
	t.query = query
	return t
}


func (t *GetNotificationBySlugTrackingNumberRequest) BuildHeader(h http.Header) *GetNotificationBySlugTrackingNumberRequest {
	t.header = h
	return t
}

func (t *GetNotificationBySlugTrackingNumberRequest) BuildPath( slug string, trackingNumber string,) *GetNotificationBySlugTrackingNumberRequest {
	t.slug = slug
	t.trackingNumber = trackingNumber
	return t
}

func (t *GetNotificationBySlugTrackingNumberRequest) isPathParamValid() error {
    if t.slug == "" {
        return errors.New("path param `slug` can not be empty")
    }
    return nil
    if t.trackingNumber == "" {
        return errors.New("path param `trackingNumber` can not be empty")
    }
    return nil
}

func (t *GetNotificationBySlugTrackingNumberRequest) build() (*http.Request, error) {
	q, err := query.Values(t.query)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
    if err := t.isPathParamValid(); err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
    }
    uri := fmt.Sprintf("/tracking/2024-04/notifications/%s/%s?",  t.slug, t.trackingNumber,)  + q.Encode()
    req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	t.requestBuilder.appendHeader(t.header, req)
	return req, nil
}

func (t *GetNotificationBySlugTrackingNumberRequest) Execute() ( *model.Notification , error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.GetNotificationBySlugTrackingNumberResponse
	err = t.sender.Do(req, &data)
	return &data.Notification,err
}