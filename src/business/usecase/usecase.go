package usecase

import (
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Interface interface{}
type usecase struct {
	cfg     config.AppConfig
	queries domain.Queries
}

func New(cfg config.AppConfig, queries domain.Queries) Interface {
	result := &usecase{
		cfg:     cfg,
		queries: queries,
	}
	return result
}
