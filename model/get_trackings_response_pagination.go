package model

type TrackingPagination struct {
    Pagination
    Tracking []Tracking `json:"trackings"`
}