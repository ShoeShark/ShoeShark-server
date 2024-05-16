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

func (rp *Repository) SaveOne(value TableNameAble) error {
	tableNameAble, ok := value.(TableNameAble)
	if !ok {
		return errors.New("value doesn't implement TableNameAble")
	}

	var err error
	if err = rp.Save(value).Error; err != nil {
		log.Error("pgsql", err, fmt.Sprintf("Failed to save %s, the value is %+v", tableNameAble.TableName(), value))
	}
	return err
}

func (rp *Repository) GetOne(result interface{}, query interface{}, args ...interface{}) (found bool, err error) {
	var (
		tableNameAble TableNameAble
		ok            bool
	)

	if tableNameAble, ok = query.(TableNameAble); !ok {
		if tableNameAble, ok = result.(TableNameAble); !ok {
			return false, errors.New("neither the query nor result implement TableNameAble")
		}
	}

	err = rp.Table(tableNameAble.TableName()).Where(query, args...).First(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("pgsql", fmt.Sprintf("record not found for query %s, the query is %+v, args are %+v", tableNameAble.TableName(), query, args))
		return false, nil
	}

	if err != nil {
		log.Error("pgsql", err, fmt.Sprintf("failed to query %s, the query is %+v, args are %+v", tableNameAble.TableName(), query, args))
		return false, err
	}

	return true, nil
}
