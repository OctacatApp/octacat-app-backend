package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/irdaislakhuafa/go-sdk/log"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/connection"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/wss"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/configreader"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/files"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/flags"
	"github.com/irdaislakhuafa/octacat-app-backend/src/middlewares"
	"github.com/rs/zerolog"
)

const (
	configPath = "./etc/cfg"
	configFile = "conf.json"
)

func main() {
	ctx := context.Background()
	log := log.Init(log.Config{Level: zerolog.LevelDebugValue})

	// check is cfg dir is exist
	if !files.IsFileExist(configPath) {
		log.Fatal(ctx, fmt.Sprintf("dir %v doesn't exist!", configPath))
		return
	}

	// parse flags for env of app
	env, err := flags.ParseFlags(configPath, configFile)
	if err != nil {
		log.Fatal(ctx, err)
	}

	// read config file
	cfg, err := configreader.ReadConfigFile(fmt.Sprintf("%v/%v/%v", configPath, *env, configFile))
	if err != nil {
		log.Fatal(ctx, err)
	}

	// init psql db
	psqlDB, err := connection.NewPostgreSQL(cfg)
	if err != nil {
		log.Fatal(ctx, err)
	}
	defer psqlDB.Close()

	// init generated code from sqlc
	gen := generated.New(psqlDB)

	// init domain
	domain := domain.New(cfg, &gen, psqlDB)

	// init usecase
	usecase := usecase.New(cfg, &domain)

	// init server
	server := http.DefaultServeMux

	// init and run websocket
	server = wss.InitAndRun(cfg, server)

	// init and run graphql server
	server = gql.InitAndRun(ctx, cfg, &usecase, server, log)

	// middlewares
	handler := middlewares.GraphQLMiddleware(cfg, &usecase)(server)

	// start server
	log.Fatal(ctx, http.ListenAndServe(":"+cfg.App.Router.GQL.Port, handler))
}
