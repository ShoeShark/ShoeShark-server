package db

type Page struct {
	Page    int         `json:"page"`
	Total   int64       `json:"total"`
	Size    int         `json:"size"`
	Records interface{} `form:"records" json:"records" uri:"records"`
}
