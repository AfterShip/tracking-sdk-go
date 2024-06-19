// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
    "errors"
	"fmt"
	"net/http"
	"github.com/google/go-querystring/query"
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/errorx"
	"github.com/AfterShip/tracking-sdk-go/v4/model"
)

type RetrackTrackingBySlugTrackingNumberRequest struct {
	slug     string
	trackingNumber     string
    query  model.RetrackTrackingBySlugTrackingNumberQuery
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewRetrackTrackingBySlugTrackingNumberRequest(sender *component.HttpSender) *RetrackTrackingBySlugTrackingNumberRequest {
	return &RetrackTrackingBySlugTrackingNumberRequest{
        sender: sender,
	}
}

func (t *RetrackTrackingBySlugTrackingNumberRequest) BuildQuery(query model.RetrackTrackingBySlugTrackingNumberQuery) *RetrackTrackingBySlugTrackingNumberRequest {
	t.query = query
	return t
}


func (t *RetrackTrackingBySlugTrackingNumberRequest) BuildHeader(h http.Header) *RetrackTrackingBySlugTrackingNumberRequest {
	t.header = h
	return t
}

func (t *RetrackTrackingBySlugTrackingNumberRequest) BuildPath( slug string, trackingNumber string,) *RetrackTrackingBySlugTrackingNumberRequest {
	t.slug = slug
	t.trackingNumber = trackingNumber
	return t
}

func (t *RetrackTrackingBySlugTrackingNumberRequest) isPathParamValid() error {
    if t.slug == "" {
        return errors.New("path param `slug` can not be empty")
    }
    return nil
    if t.trackingNumber == "" {
        return errors.New("path param `trackingNumber` can not be empty")
    }
    return nil
}

func (t *RetrackTrackingBySlugTrackingNumberRequest) build() (*http.Request, error) {
	q, err := query.Values(t.query)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
    if err := t.isPathParamValid(); err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
    }
    uri := fmt.Sprintf("/tracking/2024-04/trackings/%s/%s/retrack?",  t.slug, t.trackingNumber,)  + q.Encode()
    req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	t.requestBuilder.appendHeader(t.header, req)
	return req, nil
}

func (t *RetrackTrackingBySlugTrackingNumberRequest) Execute() ( *model.PartialUpdateTracking , error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.RetrackTrackingBySlugTrackingNumberResponse
	err = t.sender.Do(req, &data)
	return &data.Tracking,err
}