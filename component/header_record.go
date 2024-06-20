// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package component

import "time"

type RateLimit struct {
	Reset     int64 `json:"reset"`
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
}

func (r *RateLimit) isExceeded() bool {
	return r.Remaining == 0 && r.Reset >= time.Now().Unix()
}
