package model

type Pagination struct {
	Total       int    `json:"count"`
	NextCursor  string `json:"next_cursor"`
	HasNextPage bool   `json:"has_next_page"`
}
