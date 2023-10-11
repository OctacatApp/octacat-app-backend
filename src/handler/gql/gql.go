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
	"github.com/irdaislakhuafa/octacat-app-backend/src/middlewares"
)

const defaultPort = "8080"

func InitAndRun(cfg *config.AppConfig, uc *usecase.Usecase) {
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

	server := http.DefaultServeMux
	server.Handle("/", playground.Handler("GraphQL playground", "/query"))
	server.Handle("/query", srv)

	handler := (http.Handler)(server)
	handler = middlewares.GraphQLMiddleware(cfg, uc)(handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.App.Router.GQL.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.App.Router.GQL.Port, handler))
}
