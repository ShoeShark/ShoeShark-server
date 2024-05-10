package global

import (
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/repository"
)

var (
	REP    *repository.Repository
	CONFIG *config.Config
)
