package res

import (
	"github.com/shoe-shark/shoe-shark-service/mods/points/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/points/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

type PointsLogRes struct {
	Points     int64           `json:"points"`
	IsSyncLink bool            `json:"isSyncLink"`
	Source     string          `json:"source"`
	CreatedAt  util.CustomTime `json:"createdAt"`
}

// ConvertToPointsLogRes converts SstContractTransaction to PointsLogRes
func ConvertToPointsLogRes(sc *schema.UserPointsLog) PointsLogRes {
	var source = ""
	if sc.Source == string(constants.SIGN_IN) {
		source = "签到"
	} else if sc.Source == string(constants.PUBLISH_CONTENT) {
		source = "发布文章"
	}

	return PointsLogRes{
		Points:     sc.Points,
		Source:     source,
		IsSyncLink: sc.IsSyncLink,
		CreatedAt:  util.CustomTime{Time: sc.CreatedAt},
	}
}
