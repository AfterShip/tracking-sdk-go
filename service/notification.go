// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package service

import (
	"github.com/AfterShip/tracking-sdk-go/v4/component"
	"github.com/AfterShip/tracking-sdk-go/v4/request"
)

type NotificationService struct {
	sender *component.HttpSender
}

func NewNotificationService(sender *component.HttpSender) *NotificationService {
	return &NotificationService{sender: sender}
}

func (this *NotificationService) DeleteNotificationByTrackingId() *request.DeleteNotificationByTrackingIdRequest {
	return request.NewDeleteNotificationByTrackingIdRequest(this.sender)
}

func (this *NotificationService) GetNotificationByTrackingId() *request.GetNotificationByTrackingIdRequest {
	return request.NewGetNotificationByTrackingIdRequest(this.sender)
}

func (this *NotificationService) AddNotificationByTrackingId() *request.AddNotificationByTrackingIdRequest {
	return request.NewAddNotificationByTrackingIdRequest(this.sender)
}

func (this *NotificationService) AddNotificationBySlugTrackingNumber() *request.AddNotificationBySlugTrackingNumberRequest {
	return request.NewAddNotificationBySlugTrackingNumberRequest(this.sender)
}

func (this *NotificationService) GetNotificationBySlugTrackingNumber() *request.GetNotificationBySlugTrackingNumberRequest {
	return request.NewGetNotificationBySlugTrackingNumberRequest(this.sender)
}

func (this *NotificationService) DeleteNotificationBySlugTrackingNumber() *request.DeleteNotificationBySlugTrackingNumberRequest {
	return request.NewDeleteNotificationBySlugTrackingNumberRequest(this.sender)
}