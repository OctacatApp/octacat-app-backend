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

func ToUserPaginationModel(params psql.GetListUserWithPaginationParams, totalData int64, values ...psql.User) (model.UserPagination, error) {
	result := model.UserPagination{
		Limit:     int(params.Limit),
		Page:      int(params.Offset),
		TotalPage: 0,
		TotalData: int(totalData),
		Data:      []*model.User{},
	}

	for {
		if totalData > 0 {
			totalData -= int64(params.Limit)
			result.TotalPage++
		} else {
			break
		}
	}

	for _, v := range values {
		data, err := ToUserModel(v)
		if err != nil {
			return model.UserPagination{}, err
		}
		result.Data = append(result.Data, &data)
	}

	return result, nil
}
