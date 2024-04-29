package user

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/convert"
	"github.com/irdaislakhuafa/go-sdk/cryptography"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/entity"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/tokens"
)

type Interface interface {
	Register(ctx context.Context, params psql.CreateUserParams) (psql.User, error)
	Login(ctx context.Context, params psql.User) (tokens.JWTResponse, error)
	GetListWithPagination(ctx context.Context, params psql.GetListUserWithPaginationParams) ([]psql.User, error)
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (psql.User, error)
}

type user struct {
	cfg    *config.AppConfig
	domain *domain.Domain
}

func New(cfg *config.AppConfig, domain *domain.Domain) Interface {
	result := &user{
		cfg:    cfg,
		domain: domain,
	}
	return result
}

func (u *user) Register(ctx context.Context, params psql.CreateUserParams) (psql.User, error) {
	// check is email already exists? then return errror if true
	_, err := u.domain.User.GetByEmail(ctx, params.Email)
	if err != nil {
		if code := errors.GetCode(err); code.IsNotOneOf(codes.CodeSQLRecordDoesNotExist) {
			return psql.User{}, errors.NewWithCode(codes.CodeBadRequest, "error while get user by email, %v", err)
		}
	} else {
		return psql.User{}, errors.NewWithCode(codes.CodeBadRequest, "user with this email already exists")
	}

	// hash password with argon2
	if params.Password, err = cryptography.NewArgon2().Hash([]byte(params.Password)); err != nil {
		return psql.User{}, errors.NewWithCode(errors.GetCode(err), "cannot hash password, %v", err)
	}

	// create user
	user, err := u.domain.User.Create(ctx, params)
	if err != nil {
		return psql.User{}, errors.NewWithCode(errors.GetCode(err), "cannot create user, %v", err)
	}

	return user, nil
}

func (u *user) Login(ctx context.Context, params psql.User) (tokens.JWTResponse, error) {
	if params.Email == "" {
		return tokens.JWTResponse{}, errors.NewWithCode(codes.CodeBadRequest, "parameter {email} is required")
	}
	if params.Password == "" {
		return tokens.JWTResponse{}, errors.NewWithCode(codes.CodeBadRequest, "parameter {password} is required")
	}

	user, err := u.domain.User.GetByEmail(ctx, params.Email)
	if err != nil {
		if code := errors.GetCode(err); code.IsOneOf(codes.CodeSQLRecordDoesNotExist) {
			return tokens.JWTResponse{}, errors.NewWithCode(codes.CodeUnauthorized, "user with this email not registered, %v", err)
		}
		return tokens.JWTResponse{}, errors.NewWithCode(codes.CodeUnauthorized, err.Error())
	}

	if _, err := cryptography.NewArgon2().Compare([]byte(params.Password), []byte(user.Password)); err != nil {
		return tokens.JWTResponse{}, errors.NewWithCode(codes.CodeUnauthorized, "password is not match")
	}

	now := appcontext.GetRequestStartTime(ctx)

	expAt := now.Add(time.Minute * time.Duration(u.cfg.App.JWT.ExpInMinute))
	issAt := now
	claims := entity.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issAt),
			ExpiresAt: jwt.NewNumericDate(expAt),
		},
	}

	tokenString, err := tokens.NewJWT(claims, []byte(u.cfg.App.JWT.Secret))
	if err != nil {
		return tokens.JWTResponse{}, errors.NewWithCode(codes.CodeInternalServerError, err.Error())
	}

	layout := "02/01/2006 15:04:05"
	result := tokens.JWTResponse{
		Message: fmt.Sprintf("token created at '%v' and will be expired at '%v'", issAt.Format(layout), expAt.Format(layout)),
		Token:   convert.ToSafeValue[string](tokenString),
	}

	return result, nil
}

func (u *user) GetListWithPagination(ctx context.Context, params psql.GetListUserWithPaginationParams) ([]psql.User, error) {
	if params.Limit <= 0 {
		return nil, errors.NewWithCode(codes.CodeBadRequest, "minimum limit is once")
	}
	params.Offset = (params.Offset - 1) * params.Limit
	results, err := u.domain.User.GetListWithPagination(ctx, params)
	if err != nil {
		return nil, errors.NewWithCode(errors.GetCode(err), err.Error())
	}

	return results, nil
}

func (u *user) Count(ctx context.Context) (int64, error) {
	result, err := u.domain.User.Count(ctx)
	if err != nil {
		return 0, errors.NewWithCode(errors.GetCode(err), err.Error())
	}
	return result, nil
}

func (u *user) GetByID(ctx context.Context, id string) (psql.User, error) {
	result, err := u.domain.User.GetByID(ctx, id)
	if err != nil {
		return psql.User{}, errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}
	return result, nil
}
