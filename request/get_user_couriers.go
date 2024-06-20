// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
	"fmt"
	"net/http"
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/errorx"
	"github.com/AfterShip/tracking-sdk-go/v4/model"
)

type GetUserCouriersRequest struct {
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewGetUserCouriersRequest(sender *component.HttpSender) *GetUserCouriersRequest {
	return &GetUserCouriersRequest{
        sender: sender,
	}
}



func (t *GetUserCouriersRequest) BuildHeader(h http.Header) *GetUserCouriersRequest {
	t.header = h
	return t
}



func (t *GetUserCouriersRequest) build() (*http.Request, error) {
    uri := fmt.Sprintf("/tracking/2024-04/couriers?") 
    req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	t.requestBuilder.appendHeader(t.header, req)
	return req, nil
}

func (t *GetUserCouriersRequest) Execute() (  *model.GetUserCouriersResponse, error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.GetUserCouriersResponse
	return &data, t.sender.Do(req, &data)
}