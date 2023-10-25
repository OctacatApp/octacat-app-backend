package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Interface interface {
	Create(ctx context.Context, params psql.CreateUserParams) (psql.User, error)
	GetByEmail(ctx context.Context, email string) (psql.User, error)
	GetListWithPagination(ctx context.Context, params psql.GetListUserWithPaginationParams) ([]psql.User, error)
}
type user struct {
	cfg    *config.AppConfig
	gen    *generated.Generated
	psqlDB *sql.DB
}

func New(cfg *config.AppConfig, gen *generated.Generated, psqlDB *sql.DB) Interface {
	result := &user{
		cfg:    cfg,
		gen:    gen,
		psqlDB: psqlDB,
	}
	return result
}

func (u *user) Create(ctx context.Context, params psql.CreateUserParams) (psql.User, error) {
	tx, err := u.psqlDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return psql.User{}, err
	}
	defer tx.Rollback()

	params.ID = uuid.NewString()
	params.CreatedAt = time.Now()
	params.CreatedBy = u.cfg.App.Default.Me

	qtx := u.gen.PSQL.WithTx(tx)
	user, err := qtx.CreateUser(ctx, params)
	if err != nil {
		return psql.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return psql.User{}, err
	}

	return user, nil
}

func (u *user) GetByEmail(ctx context.Context, email string) (psql.User, error) {
	user, err := u.gen.PSQL.GetUserByEmail(ctx, email)
	if err != nil {
		return psql.User{}, err
	}

	return user, nil
}

func (u *user) GetListWithPagination(ctx context.Context, params psql.GetListUserWithPaginationParams) ([]psql.User, error) {
	users, err := u.gen.PSQL.GetListUserWithPagination(ctx, params)
	if err != nil {
		return nil, err
	}
	return users, nil
}
