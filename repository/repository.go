package repository

import (
	"errors"
	"fmt"
	"github.com/shoe-shark/shoe-shark-service/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TableNameAble interface {
	TableName() string
}

type Repository struct {
	*gorm.DB
}

func (dw *Repository) SaveOne(value TableNameAble) error {
	tableNameAble, ok := value.(TableNameAble)
	if !ok {
		return errors.New("value doesn't implement TableNameAble")
	}

	var err error
	if err = dw.Save(value).Error; err != nil {
		log.Error("pg", err, fmt.Sprintf("Failed to save %s, the value is %+v", tableNameAble.TableName(), value))
	}
	return err
}

func InitDb(config *config.Config) *Repository {
	var dw Repository
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Database.Pg.Host,
		config.Database.Pg.Username,
		config.Database.Pg.Password,
		config.Database.Pg.Dbname,
		config.Database.Pg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dw = Repository{db}
	return &dw
}
