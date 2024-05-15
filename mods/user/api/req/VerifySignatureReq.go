package req

type VerifySignatureReq struct {
	Signature      string `form:"signature" json:"signature" binding:"required"`
	AccountAddress string `form:"accountAddress" json:"accountAddress" binding:"required"`
}
