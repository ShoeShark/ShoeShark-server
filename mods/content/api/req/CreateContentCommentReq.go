package req

type CreateContentCommentReq struct {
	Description string `json:"description"`
	ContentId   string `json:"contentId"`
}
