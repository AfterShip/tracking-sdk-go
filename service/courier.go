// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package service

import (
    "github.com/AfterShip/tracking-sdk-go/v4/component"
    "github.com/AfterShip/tracking-sdk-go/v4/request"
)

type CourierService struct {
    sender *component.HttpSender
}

func NewCourierService(sender *component.HttpSender) *CourierService {
    return &CourierService{sender: sender}
}

func (this *CourierService) DetectCourier() *request.DetectCourierRequest {
    return request.NewDetectCourierRequest(this.sender)
}

func (this *CourierService) GetUserCouriers() *request.GetUserCouriersRequest {
    return request.NewGetUserCouriersRequest(this.sender)
}

func (this *CourierService) GetAllCouriers() *request.GetAllCouriersRequest {
    return request.NewGetAllCouriersRequest(this.sender)
}

