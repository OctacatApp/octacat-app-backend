package generated

import "github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"

type Generated struct {
	PSQL *psql.Queries
}

func New(psqlDB psql.DBTX) Generated {
	result := Generated{
		PSQL: psql.New(psqlDB),
	}
	return result
}
