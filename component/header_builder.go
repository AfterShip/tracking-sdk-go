// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package component

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type headerBuilder struct {
	header http.Header
	auth   *Authenticator
	err    error
}

const defaultUserAgent = "tracking-sdk-go/v4 (https://www.aftership.com) Go-http-client/1.1"

func newHeaderBuilder(auth *Authenticator) *headerBuilder {
	return &headerBuilder{header: http.Header{}, auth: auth}
}

func (h *headerBuilder) buildDefaultHeader(userAgent string) *headerBuilder {
	h.header.Add("content-type", "application/json")
	h.header.Add("date", time.Now().UTC().Format(http.TimeFormat))
	h.header.Add("request-id", uuid.New().String())
	h.header.Add("user-agent", userAgent)
	h.header.Add("aftership-client", defaultUserAgent)
	return h
}

func (h *headerBuilder) buildAuthHeader(req *http.Request) *headerBuilder {
	for k, v := range req.Header {
		if len(v) > 0 {
			h.header.Set(k, v[0])
		}
	}
	req.Header = h.header
	m, err := h.auth.Sign(req)
	if err != nil {
		h.err = err
		return h
	}
	for key, val := range m {
		h.header.Set(key, val)
	}
	return h
}

func (h *headerBuilder) build() (http.Header, error) {
	return h.header, h.err
}
