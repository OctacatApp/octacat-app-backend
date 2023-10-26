package directives

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/entity"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/key"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/tokens"
)

func JWTDirective(ctx context.Context, _ any, next graphql.Resolver) (res any, err error) {
	headers := http.Header{}
	if h, isOk := ctx.Value(key.Key("headers")).(http.Header); isOk {
		headers = h
	}

	authorization := headers.Get("authorization")
	mustPrefix := "bearer"

	if len(strings.Trim(authorization, " ")) <= len(mustPrefix) {
		return nil, errors.New(fmt.Sprintf("the 'authorization' header is empty or not valid"))
	}

	if prefix := authorization[0:len(mustPrefix)]; !strings.EqualFold(prefix, mustPrefix) {
		return nil, errors.New(fmt.Sprintf("your 'authorization' header not start with '%v'", mustPrefix))
	}

	cfg := ctx.Value(key.Key("cfg")).(*config.AppConfig)
	tokenString := authorization[len(mustPrefix)+1:]
	token, err := tokens.Validate(tokenString, []byte(cfg.App.JWT.Secret), &entity.Claims{})
	if err != nil {
		return nil, err
	}

	claims, err := tokens.GetClaims[*entity.Claims](token)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, key.Key("claims"), *claims)
	return next(ctx)
}
