package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated/psql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

type Interface interface {
	Create(ctx context.Context, params psql.CreateUserParams) (psql.User, error)
	GetByEmail(ctx context.Context, email string) (psql.User, error)
	GetListWithPagination(ctx context.Context, params psql.GetListUserWithPaginationParams) ([]psql.User, error)
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (psql.User, error)
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
		return psql.User{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	params.ID = uuid.NewString()
	params.CreatedAt = appcontext.GetRequestStartTime(ctx)
	params.CreatedBy = u.cfg.App.Default.Me

	qtx := u.gen.PSQL.WithTx(tx)
	user, err := qtx.CreateUser(ctx, params)
	if err != nil {
		return psql.User{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if err := tx.Commit(); err != nil {
		return psql.User{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return user, nil
}

func (u *user) GetByEmail(ctx context.Context, email string) (psql.User, error) {
	user, err := u.gen.PSQL.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return psql.User{}, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
		}
		return psql.User{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return user, nil
}

func (u *user) GetListWithPagination(ctx context.Context, params psql.GetListUserWithPaginationParams) ([]psql.User, error) {
	users, err := u.gen.PSQL.GetListUserWithPagination(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
		}
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}
	return users, nil
}

func (u *user) Count(ctx context.Context) (int64, error) {
	total, err := u.gen.PSQL.CountUser(ctx)
	if err != nil {
		return 0, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return total, nil
}

func (u *user) GetByID(ctx context.Context, id string) (psql.User, error) {
	result, err := u.gen.PSQL.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return psql.User{}, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
		}
		return psql.User{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return result, nil
}
