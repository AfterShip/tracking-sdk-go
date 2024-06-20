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

type DeleteNotificationByTrackingIdRequest struct {
	trackingId     string
	body  model.DeleteNotificationByTrackingIdRequest
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewDeleteNotificationByTrackingIdRequest(sender *component.HttpSender) *DeleteNotificationByTrackingIdRequest {
	return &DeleteNotificationByTrackingIdRequest{
        sender: sender,
	}
}


func (t *DeleteNotificationByTrackingIdRequest) BuildBody(body model.DeleteNotificationByTrackingIdRequest) *DeleteNotificationByTrackingIdRequest {
	t.body = body
	return t
}

func (t *DeleteNotificationByTrackingIdRequest) BuildHeader(h http.Header) *DeleteNotificationByTrackingIdRequest {
	t.header = h
	return t
}

func (t *DeleteNotificationByTrackingIdRequest) BuildPath( trackingId string,) *DeleteNotificationByTrackingIdRequest {
	t.trackingId = trackingId
	return t
}

func (t *DeleteNotificationByTrackingIdRequest) isPathParamValid() error {
    if t.trackingId == "" {
        return errors.New("path param `trackingId` can not be empty")
    }
    return nil
}

func (t *DeleteNotificationByTrackingIdRequest) build() (*http.Request, error) {
    if err := t.isPathParamValid(); err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
    }
    uri := fmt.Sprintf("/tracking/2024-04/notifications/%s/remove?",  t.trackingId,) 
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

func (t *DeleteNotificationByTrackingIdRequest) Execute() ( *model.Notification , error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.DeleteNotificationByTrackingIdResponse
	err = t.sender.Do(req, &data)
	return &data.Notification,err
}