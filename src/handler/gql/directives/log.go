package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/log"
)

func LogDirective(log log.Interface) func(ctx context.Context, _ any, next graphql.Resolver) (res any, err error) {
	return func(ctx context.Context, _ any, next graphql.Resolver) (res any, err error) {
		res, err = next(ctx)
		code := errors.GetCode(err)
		if code.IsNotOneOf(codes.NoCode) {
			log.Error(ctx, err)
		}
		return
	}
}
