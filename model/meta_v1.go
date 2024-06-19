// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package model

// MetaV1 Meta data
type MetaV1 struct {
    // Code meta code
    Code int `json:"code"`
    // Message error message, only exist if the response status is not 2xx
    Message string `json:"message,omitempty"`
    // Type error type, only exist if the response status is not 2xx
    Type string `json:"type,omitempty"`
}

