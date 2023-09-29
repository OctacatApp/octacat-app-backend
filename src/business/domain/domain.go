package domain

import (
	"database/sql"

	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain/user"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Domain struct {
	User user.Interface
}

func New(cfg *config.AppConfig, gen *generated.Generated, psqlDB *sql.DB) Domain {
	result := Domain{
		User: user.New(cfg, gen, psqlDB),
	}
	return result
}
