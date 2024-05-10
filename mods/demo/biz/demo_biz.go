package biz

import (
	"context"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/global"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/dao"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/schema"
)

func SaveOne(ctx context.Context, item *schema.DemoForm) error {
	demo := &schema.Demo{
		Name: item.Name,
		Code: item.Code,
	}

	err := dao.SaveOne(global.REP, demo)
	if err != nil {
		return errors.New("save demo error")
	}
	return nil
}
