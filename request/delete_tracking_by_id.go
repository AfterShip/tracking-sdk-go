// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
    "errors"
	"fmt"
	"net/http"
	"github.com/aftership/tracking-sdk-go/v4/component"
	"github.com/aftership/tracking-sdk-go/v4/errorx"
	"github.com/aftership/tracking-sdk-go/v4/model"
)

type DeleteTrackingByIdRequest struct {
	id     string
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewDeleteTrackingByIdRequest(sender *component.HttpSender) *DeleteTrackingByIdRequest {
	return &DeleteTrackingByIdRequest{
        sender: sender,
	}
}



func (t *DeleteTrackingByIdRequest) BuildHeader(h http.Header) *DeleteTrackingByIdRequest {
	t.header = h
	return t
}

func (t *DeleteTrackingByIdRequest) BuildPath( id string,) *DeleteTrackingByIdRequest {
	t.id = id
	return t
}

func (t *DeleteTrackingByIdRequest) isPathParamValid() error {
    if t.id == "" {
        return errors.New("path param `id` can not be empty")
    }
    return nil
}

func (t *DeleteTrackingByIdRequest) build() (*http.Request, error) {
    if err := t.isPathParamValid(); err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
    }
    uri := fmt.Sprintf("/tracking/2024-04/trackings/%s?",  t.id,) 
    req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	t.requestBuilder.appendHeader(t.header, req)
	return req, nil
}

func (t *DeleteTrackingByIdRequest) Execute() ( *model.PartialDeleteTracking , error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.DeleteTrackingByIdResponse
	err = t.sender.Do(req, &data)
	return &data.Tracking,err
}