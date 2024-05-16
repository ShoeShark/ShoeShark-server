package req

type QueryContentReq struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	AccountAddress string `json:"accountAddress"`
	Page           int    `json:"page"`
	Size           int    `json:"size"`
}
