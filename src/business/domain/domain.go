package domain

import "github.com/irdaislakhuafa/octacat-app-backend/src/business/domain/psql"

type Domain struct {
	PSQL *psql.Queries
}

func New(psqlDB psql.DBTX) Domain {
	result := Domain{
		PSQL: psql.New(psqlDB),
	}
	return result
}
