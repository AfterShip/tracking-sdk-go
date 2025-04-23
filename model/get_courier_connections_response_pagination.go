package model

type CourierConnectionPagination struct {
	Pagination
	CourierConnection []CourierConnection `json:"courier_connections"`
}
