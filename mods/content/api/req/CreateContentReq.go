package req

type CreateContentReq struct {
	Title          string `json:"title" binding:"required"`
	Body           string `json:"body" binding:"required"`
	AccountAddress string `json:"accountAddress" binding:"required"`
	Location       string `json:"location"`
	IsPublic       bool   `json:"isPublic"`
}
