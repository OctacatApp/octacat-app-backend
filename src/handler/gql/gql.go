package gql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/directives"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/generated/server"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

const defaultPort = "8080"

func InitAndRun(cfg *config.AppConfig, uc *usecase.Usecase, serverMux *http.ServeMux) *http.ServeMux {
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.App.Router.GQL.Port)
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
