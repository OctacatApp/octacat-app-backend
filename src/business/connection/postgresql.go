package connection

import (
	"database/sql"
	"fmt"

	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/operator"
	_ "github.com/lib/pq"
)

func NewPostgreSQL(cfg *config.AppConfig) *sql.DB {
	psql := cfg.DataSource.PostgreSQL
	datasource := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", psql.Host, psql.Port, psql.Username, psql.Password, psql.Database, operator.Ternary(psql.SSL, "enable", "disable"))
	db, err := sql.Open(psql.Driver, datasource)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to %v database, %v", psql.Driver, err))
	}

	return db
}
