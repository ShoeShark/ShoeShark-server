package biz

import (
	"context"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/mods/points/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/points/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/points/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/points/dao"
	"github.com/shoe-shark/shoe-shark-service/mods/points/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	"time"
)

func SignIn(ctx *context.Context) error {
	accountAddress := (*ctx).Value("accountAddress").(string)

	isSignIn, err := dao.VerifyIsSignIn(db.GetPGRepository(), accountAddress)
	if err != nil {
		return err
	}

	if isSignIn {
		return errors.New("已签到")
	}

	insertLog := schema.UserPointsLog{
		AccountAddress: accountAddress,
		Points:         5,
		IsSyncLink:     false,
		Source:         string(constants.SIGN_IN),
		CreatedAt:      time.Now(),
	}

	return dao.AddPoint(db.GetPGRepository(), &insertLog)
}

func GetPointsLogs(ctx *context.Context, queryReq *req.QueryPointsLogReq) (*db.Page, error) {
	// 构建查询条件
	accountAddress := (*ctx).Value("accountAddress").(string)
	rp := repository.GetPGRepository()
	dbQuery := rp.Model(&schema.UserPointsLog{}).Where("account_address = ?", accountAddress)

	var total int64
	var records []schema.UserPointsLog
	var page = &db.Page{
		Page:    queryReq.Page,
		Size:    queryReq.Size,
		Records: &records,
		Total:   total,
	}

	dbQuery.Count(&total)
	if total < 1 {
		return page, nil
	}

	offset := (queryReq.Page - 1) * queryReq.Size
	err := dbQuery.Offset(offset).Limit(queryReq.Size).Find(&records).Error
	if err != nil {
		return nil, err
	}

	var pointsLogInfo []res.PointsLogRes
	err = util.GenericConvert(&records, &pointsLogInfo)

	for i, logRes := range pointsLogInfo {
		if logRes.Source == string(constants.SIGN_IN) {
			pointsLogInfo[i].Source = "签到"
		} else if logRes.Source == string(constants.PUBLISH_CONTENT) {
			pointsLogInfo[i].Source = "发布文章"
		}
	}

	return page, nil
}