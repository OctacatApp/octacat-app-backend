package transform

import (
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/generated/model"
)

func FromGetListParams(params model.GetListParams) (psql.GetListUserWithPaginationParams, error) {
	result := psql.GetListUserWithPaginationParams{
		Limit:  int32(params.Limit),
		Offset: int32(params.Page),
	}
	return result, nil
}
