// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package model

// NotificationResponseV1 
type NotificationResponseV1 struct {
    // Meta Meta data
    Meta MetaV1 `json:"meta,omitempty"`
    // Data 
    Data *DataNotificationResponseV1 `json:"data,omitempty"`
}

// DataNotificationResponseV1 
type DataNotificationResponseV1 struct {
    // Notification Object describes the notification information.
    Notification Notification `json:"notification,omitempty"`
}

