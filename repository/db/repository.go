package db

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
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
