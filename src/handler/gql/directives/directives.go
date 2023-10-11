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
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/tokens"
)

func JWTDirective(ctx context.Context, _ any, next graphql.Resolver) (res any, err error) {
	headers := ctx.Value("headers").(http.Header)
	authorization := headers.Get("authorization")
	basic := "Basic"

	if len(strings.Trim(authorization, " ")) <= len(basic) {
		return nil, errors.New(fmt.Sprintf("the 'authorization' header is empty or not valid"))
	}

	if prefix := authorization[0:len(basic)]; !strings.EqualFold(prefix, basic) {
		return nil, errors.New(fmt.Sprintf("your 'authorization' header not start with '%v'", basic))
	}

	cfg := ctx.Value("cfg").(*config.AppConfig)
	tokenString := authorization[len(basic)+1:]
	token, err := tokens.Validate(tokenString, []byte(cfg.App.JWT.Secret), entity.Claims{})
	if err != nil {
		return nil, err
	}
	// fmt.Printf("token: '%v\n", token)

	claims, err := tokens.GetClaims[entity.Claims](token)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "claims", claims)
	return next(ctx)
}
