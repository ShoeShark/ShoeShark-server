package db

import (
	"fmt"
	"github.com/shoe-shark/shoe-shark-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgRw Repository

func Init(config *config.Config) {
	InitPg(config.Database.Pg)
}

func GetPGRepository() *Repository {
	return &pgRw
}

func InitPg(config *config.DatabasePg) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.Username,
		config.Password,
		config.Dbname,
		config.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	pgRw = Repository{db}
}
