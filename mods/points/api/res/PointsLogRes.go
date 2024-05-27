package res

import "time"

type PointsLogRes struct {
	Points     int64     `json:"points"`
	IsSyncLink bool      `json:"isSyncLink"`
	Source     string    `json:"source"`
	CreatedAt  time.Time `json:"createdAt"`
}
