package biz

import (
	"context"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/mods/points/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/points/dao"
	"github.com/shoe-shark/shoe-shark-service/mods/points/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	"time"
)

func SignIn(ctx *context.Context) error {
	accountAddress := (*ctx).Value("accountAddress").(string)

	return AddPoints(accountAddress, 5, constants.SIGN_IN)
}

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

func AddPoints(accountAddress string, points int64, addType constants.PointsSourceType) error {
	if addType == constants.SIGN_IN {
		isSignIn, err := dao.VerifyIsSignIn(db.GetPGRepository(), accountAddress)
		if err != nil {
			return err
		}

		if isSignIn {
			return errors.New("已签到")
		}
	} else if addType == constants.PUBLISH_CONTENT {
		isFull, err := dao.VerifyPublishContent(db.GetPGRepository(), accountAddress)
		if err != nil {
			return err
		}
		if isFull {
			return errors.New("发布文章获取积分已上限")
		}
	}

	insertLog := schema.UserPointsLog{
		AccountAddress: accountAddress,
		Points:         points,
		IsSyncLink:     false,
		Source:         string(constants.SIGN_IN),
		CreatedAt:      time.Now(),
	}

	err := dao.AddPointsLog(db.GetPGRepository(), &insertLog)
	if err != nil {
		return err
	}

	return dao.AddPoints(db.GetPGRepository(), accountAddress, points)
}
