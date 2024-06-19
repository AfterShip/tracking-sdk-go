// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package service

import (
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/request"
)

type TrackingService struct {
	sender *component.HttpSender
}

func NewTrackingService(sender *component.HttpSender) *TrackingService {
	return &TrackingService{sender: sender}
}

func (this *TrackingService) MarkTrackingCompletedById() *request.MarkTrackingCompletedByIdRequest {
	return request.NewMarkTrackingCompletedByIdRequest(this.sender)
}

func (this *TrackingService) DeleteTrackingBySlugTrackingNumber() *request.DeleteTrackingBySlugTrackingNumberRequest {
	return request.NewDeleteTrackingBySlugTrackingNumberRequest(this.sender)
}

func (this *TrackingService) MarkTrackingCompletedBySlugTrackingNumber() *request.MarkTrackingCompletedBySlugTrackingNumberRequest {
	return request.NewMarkTrackingCompletedBySlugTrackingNumberRequest(this.sender)
}

func (this *TrackingService) CreateTracking() *request.CreateTrackingRequest {
	return request.NewCreateTrackingRequest(this.sender)
}

func (this *TrackingService) UpdateTrackingById() *request.UpdateTrackingByIdRequest {
	return request.NewUpdateTrackingByIdRequest(this.sender)
}

func (this *TrackingService) GetTrackingBySlugTrackingNumber() *request.GetTrackingBySlugTrackingNumberRequest {
	return request.NewGetTrackingBySlugTrackingNumberRequest(this.sender)
}

func (this *TrackingService) RetrackTrackingBySlugTrackingNumber() *request.RetrackTrackingBySlugTrackingNumberRequest {
	return request.NewRetrackTrackingBySlugTrackingNumberRequest(this.sender)
}

func (this *TrackingService) GetTrackings() *request.GetTrackingsRequest {
	return request.NewGetTrackingsRequest(this.sender)
}

func (this *TrackingService) UpdateTrackingBySlugTrackingNumber() *request.UpdateTrackingBySlugTrackingNumberRequest {
	return request.NewUpdateTrackingBySlugTrackingNumberRequest(this.sender)
}

func (this *TrackingService) GetTrackingById() *request.GetTrackingByIdRequest {
	return request.NewGetTrackingByIdRequest(this.sender)
}

func (this *TrackingService) DeleteTrackingById() *request.DeleteTrackingByIdRequest {
	return request.NewDeleteTrackingByIdRequest(this.sender)
}

func (this *TrackingService) RetrackTrackingById() *request.RetrackTrackingByIdRequest {
	return request.NewRetrackTrackingByIdRequest(this.sender)
}