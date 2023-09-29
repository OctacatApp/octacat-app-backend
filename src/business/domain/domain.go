package domain

import "github.com/irdaislakhuafa/octacat-app-backend/src/business/domain/psql"

type Domain struct {
	psql *psql.Queries
}

func New(psqlDB psql.DBTX) Domain {
	result := Domain{
		psql: psql.New(psqlDB),
	}
	return result
}
