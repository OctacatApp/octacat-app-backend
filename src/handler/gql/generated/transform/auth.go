package transform

import (
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/generated/model"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/tokens"
)

func FromRegisterParams(params model.RegisterParam) (psql.CreateUserParams, error) {
	result := psql.CreateUserParams{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}
	return result, nil
}

func ToUserModel(params psql.User) (model.User, error) {
	layout := "02/01/2006 15:04:05"
	result := model.User{
		ID:           params.ID,
		Name:         params.Name,
		Email:        params.Email,
		Password:     params.Password,
		ProfileImage: params.ProfileImage,
		CreatedAt:    params.CreatedAt.Format(layout),
		CreatedBy:    params.CreatedBy,
		UpdatedAt:    params.UpdatedAt.Time.Format(layout),
		UpdatedBy:    params.UpdatedBy.String,
		DeletedAt:    params.DeletedAt.Time.Format(layout),
		DeletedBy:    params.DeletedBy.String,
		IsDeleted:    params.IsDeleted,
	}
	return result, nil
}

func FromLoginParam(params model.LoginParam) (psql.User, error) {
	result := psql.User{
		Email:    params.Email,
		Password: params.Password,
	}
	return result, nil
}

func ToJWTResponseModel(params tokens.JWTResponse) (model.JWTResponse, error) {
	result := model.JWTResponse{
		Message: params.Message,
		Token:   params.Token,
	}
	return result, nil
}
