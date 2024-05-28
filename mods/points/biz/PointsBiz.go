package biz

import (
	"context"
	"github.com/shoe-shark/shoe-shark-service/mods/points/dao"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func GetPoints(ctx *context.Context) (int64, error) {
	accountAddress := (*ctx).Value("accountAddress").(string)

	points, err := dao.GetPoints(db.GetPGRepository(), accountAddress)
	if err != nil {
		return 0, err
	}

	if points == nil {
		return 0, nil
	}

	return points.Points, nil
}
