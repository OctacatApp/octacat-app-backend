package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/irdaislakhuafa/go-argon2/argon2"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/tokens"
)

type Interface interface {
	Register(ctx context.Context, params psql.CreateUserParams) (psql.User, error)
	Login(ctx context.Context, params psql.User) (tokens.JWTResponse, error)
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
		if !errors.Is(err, sql.ErrNoRows) {
			return psql.User{}, errors.Join(errors.New("error while get user by email"), err)
		}
	} else {
		return psql.User{}, errors.New("user with this email already exists")
	}

	// hash password with argon2
	if params.Password, err = argon2.HashArgon2([]byte(params.Password)); err != nil {
		return psql.User{}, errors.Join(errors.New("cannot hash password"), err)
	}

	// create user
	user, err := u.domain.User.Create(ctx, params)
	if err != nil {
		return psql.User{}, errors.Join(errors.New("cannot create user"), err)
	}

	return user, nil
}

func (u *user) Login(ctx context.Context, params psql.User) (tokens.JWTResponse, error) {
	if params.Email == "" {
		return tokens.JWTResponse{}, errors.New("parameter {email} is required")
	}
	if params.Password == "" {
		return tokens.JWTResponse{}, errors.New("parameter {password} is required")
	}

	user, err := u.domain.User.GetByEmail(ctx, params.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tokens.JWTResponse{}, errors.Join(errors.New("user with this email not registered"), err)
		}
		return tokens.JWTResponse{}, err
	}

	if _, err := argon2.CompareArgon2(params.Password, user.Password); err != nil {
		return tokens.JWTResponse{}, errors.Join(errors.New("password is not match"), err)
	}

	expAt := time.Now().Add(time.Minute * time.Duration(u.cfg.App.JWT.ExpInMinute))
	issAt := time.Now()
	claims := tokens.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issAt),
			ExpiresAt: jwt.NewNumericDate(expAt),
		},
	}

	tokenString, err := tokens.NewJWT(claims, []byte(u.cfg.App.JWT.Secret))
	if err != nil {
		return tokens.JWTResponse{}, err
	}

	layout := "02/01/2006 15:04:05"
	result := tokens.JWTResponse{
		Message: fmt.Sprintf("token created at '%v' and will be expired at '%v'", issAt.Format(layout), expAt.Format(layout)),
		Token:   *tokenString,
	}

	return result, nil
}
