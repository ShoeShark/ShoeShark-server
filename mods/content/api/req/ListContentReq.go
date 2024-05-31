package req

type QueryContentReq struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	AccountAddress string `json:"accountAddress"`
	Location       string `json:"location"`
	Page           int    `json:"page"`
	Size           int    `json:"size"`
}
