package gql

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/go-sdk/log"
	"github.com/irdaislakhuafa/go-sdk/strformat"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/directives"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/generated/server"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

func InitAndRun(ctx context.Context, cfg *config.AppConfig, uc *usecase.Usecase, serverMux *http.ServeMux, log log.Interface) *http.ServeMux {
	srv := handler.NewDefaultServer(
		server.NewExecutableSchema(
			server.Config{
				Resolvers: &Resolver{
					Usecase: uc,
				},
				Directives: server.DirectiveRoot{
					Jwt: directives.JWTDirective,
				},
			},
		),
	)

	serverMux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	serverMux.Handle("/query", CORSHandler(srv))

	log.Info(ctx, strformat.TmplWithoutErr("connect to http://localhost:{{.Port}}/ for GraphQL playground", cfg.App.Router.GQL))
	return serverMux
}

func CORSHandler(handler http.Handler) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		cors := map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		}

		for k, v := range cors {
			w.Header().Set(k, v)
		}
		handler.ServeHTTP(w, r)
	}
	return f
}
