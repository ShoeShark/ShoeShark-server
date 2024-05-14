package biz

import (
	"context"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/schema"
	"github.com/shoe-shark/shoe-shark-service/repository"
)

func CreateDemo(ctx context.Context, item *req.DemoCreateReq) error {
	demo := &schema.Demo{
		Name: item.Name,
		Code: item.Code,
	}

	rp := repository.GetPGRepository()

	err := rp.SaveOne(demo)
	if err != nil {
		return errors.New("save demo error")
	}
	return nil
}
