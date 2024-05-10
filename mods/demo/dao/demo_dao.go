package dao

import (
	"github.com/shoe-shark/shoe-shark-service/mods/demo/schema"
	"github.com/shoe-shark/shoe-shark-service/repository"
)

func SaveOne(dw *repository.Repository, item *schema.Demo) error {
	err := dw.SaveOne(item)
	if err != nil {
		return err
	}
	return nil
}
