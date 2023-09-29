package usecase

import (
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase/user"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Usecase struct {
	User user.Interface
}

func New(cfg *config.AppConfig, domain *domain.Domain) Usecase {
	result := Usecase{
		User: user.New(cfg, domain),
	}
	return result
}
