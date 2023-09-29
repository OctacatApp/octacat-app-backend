package user

import (
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Interface interface {
}

type user struct {
	cfg    config.AppConfig
	domain domain.Domain
}

func New(cfg config.AppConfig, domain domain.Domain) Interface {
	result := &user{
		cfg:    cfg,
		domain: domain,
	}
	return result
}
