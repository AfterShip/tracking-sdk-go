// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package request

import (
	"fmt"
	"net/http"
    "bytes"
    "encoding/json"
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/errorx"
	"github.com/AfterShip/tracking-sdk-go/v4/model"
)

type PredictBatchRequest struct {
	body  model.PredictBatchRequest
	header http.Header
	sender *component.HttpSender
	requestBuilder
}

func NewPredictBatchRequest(sender *component.HttpSender) *PredictBatchRequest {
	return &PredictBatchRequest{
        sender: sender,
	}
}


func (t *PredictBatchRequest) BuildBody(body model.PredictBatchRequest) *PredictBatchRequest {
	t.body = body
	return t
}

func (t *PredictBatchRequest) BuildHeader(h http.Header) *PredictBatchRequest {
	t.header = h
	return t
}



func (t *PredictBatchRequest) build() (*http.Request, error) {
    uri := fmt.Sprintf("/tracking/2024-04/estimated-delivery-date/predict-batch?") 
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

func (t *PredictBatchRequest) Execute() (  *model.PredictBatchResponse, error) {
	req, err := t.build()
	if err != nil {
        return nil, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
	}
	var data model.PredictBatchResponse
	return &data, t.sender.Do(req, &data)
}