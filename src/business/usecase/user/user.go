package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/irdaislakhuafa/go-argon2/argon2"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Interface interface {
	Register(ctx context.Context, params psql.CreateUserParams) (psql.User, error)
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
