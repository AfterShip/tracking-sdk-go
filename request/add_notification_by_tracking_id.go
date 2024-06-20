// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
    "errors"
	"fmt"
	"net/http"
    "bytes"
    "encoding/json"
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/errorx"
	"github.com/AfterShip/tracking-sdk-go/v4/model"
)

type AddNotificationByTrackingIdRequest struct {
	trackingId     string
	body  model.AddNotificationByTrackingIdRequest
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewAddNotificationByTrackingIdRequest(sender *component.HttpSender) *AddNotificationByTrackingIdRequest {
	return &AddNotificationByTrackingIdRequest{
        sender: sender,
	}
}


func (t *AddNotificationByTrackingIdRequest) BuildBody(body model.AddNotificationByTrackingIdRequest) *AddNotificationByTrackingIdRequest {
	t.body = body
	return t
}

func (t *AddNotificationByTrackingIdRequest) BuildHeader(h http.Header) *AddNotificationByTrackingIdRequest {
	t.header = h
	return t
}

func (t *AddNotificationByTrackingIdRequest) BuildPath( trackingId string,) *AddNotificationByTrackingIdRequest {
	t.trackingId = trackingId
	return t
}

func (t *AddNotificationByTrackingIdRequest) isPathParamValid() error {
    if t.trackingId == "" {
        return errors.New("path param `trackingId` can not be empty")
    }
    return nil
}

func (t *AddNotificationByTrackingIdRequest) build() (*http.Request, error) {
    if err := t.isPathParamValid(); err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
    }
    uri := fmt.Sprintf("/tracking/2024-04/notifications/%s/add?",  t.trackingId,) 
    body,err := json.Marshal(t.body)
    if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
    }
    req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	t.requestBuilder.appendHeader(t.header, req)
	return req, nil
}

func (t *AddNotificationByTrackingIdRequest) Execute() ( *model.Notification , error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.AddNotificationByTrackingIdResponse
	err = t.sender.Do(req, &data)
	return &data.Notification,err
}