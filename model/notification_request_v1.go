// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package model

// NotificationRequestV1 Notification request object
type NotificationRequestV1 struct {
    // Emails Email address(es) to receive email notifications.Accept either array or comma separated as input.
    Emails []string `json:"emails,omitempty"`
    // Smses The phone number(s) to receive sms notifications.  Enter `+` and `area code` before phone number.Accept either array or comma separated as input.
    Smses []string `json:"smses,omitempty"`
}


