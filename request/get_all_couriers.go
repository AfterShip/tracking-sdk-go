// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
	"fmt"
	"net/http"
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/errorx"
	"github.com/AfterShip/tracking-sdk-go/v4/model"
)

type GetAllCouriersRequest struct {
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewGetAllCouriersRequest(sender *component.HttpSender) *GetAllCouriersRequest {
	return &GetAllCouriersRequest{
        sender: sender,
	}
}



func (t *GetAllCouriersRequest) BuildHeader(h http.Header) *GetAllCouriersRequest {
	t.header = h
	return t
}



func (t *GetAllCouriersRequest) build() (*http.Request, error) {
    uri := fmt.Sprintf("/tracking/2024-04/couriers/all?") 
    req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	t.requestBuilder.appendHeader(t.header, req)
	return req, nil
}

func (t *GetAllCouriersRequest) Execute() (  *model.GetAllCouriersResponse, error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.GetAllCouriersResponse
	return &data, t.sender.Do(req, &data)
}