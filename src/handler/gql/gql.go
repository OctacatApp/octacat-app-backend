package gql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql/generated/server"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/middlewares"
)

const defaultPort = "8080"

func InitAndRun(cfg *config.AppConfig, uc *usecase.Usecase, serverMux *http.ServeMux) http.Handler {
	srv := handler.NewDefaultServer(
		server.NewExecutableSchema(
			server.Config{
				Resolvers: &Resolver{
					Usecase: uc,
				},
			},
		),
	)

	// uri handler for GraphQL
	serverMux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	serverMux.Handle("/query", srv)

	// middlewares
	handler := (http.Handler)(serverMux)
	handler = middlewares.GraphQLMiddleware(cfg, uc)(handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.App.Router.GQL.Port)
	return handler
}
