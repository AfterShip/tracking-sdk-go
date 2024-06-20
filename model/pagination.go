package model

type Pagination struct {
    Total int `json:"count"`
    Page  int `json:"page"`
    Limit int `json:"limit"`
}