package repository

import (
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func GetPGRepository() *db.Repository {
	return db.GetPGRepository()
}

func Init() {
	cfg := config.GetConfig()
	db.Init(cfg)
}
