package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/irdaislakhuafa/octacat-app-backend/src/business/connection"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/generated"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/wss"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/configreader"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/files"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/flags"
)

const (
	configPath = "./etc/cfg"
	configFile = "conf.json"
)

func main() {
	// check is cfg dir is exist
	if !files.IsFileExist(configPath) {
		fmt.Printf("dir %v doesn't exist!", configPath)
		return
	}

	// parse flags for env of app
	env, err := flags.ParseFlags(configPath, configFile)
	if err != nil {
		panic(err)
	}

	// read config file
	cfg, err := configreader.ReadConfigFile(fmt.Sprintf("%v/%v/%v", configPath, *env, configFile))
	if err != nil {
		panic(err)
	}

	// init psql db
	psqlDB := connection.NewPostgreSQL(cfg)
	defer psqlDB.Close()

	// init generated code from sqlc
	gen := generated.New(psqlDB)

	// init domain
	domain := domain.New(cfg, &gen, psqlDB)

	// init usecase
	usecase := usecase.New(cfg, &domain)

	// init server
	server := http.DefaultServeMux

	// init and run graphql server
	handler := gql.InitAndRun(cfg, &usecase, server)

	// init and run websocket
	handler = wss.InitAndRun(cfg, server)

	// start server
	log.Fatal(http.ListenAndServe(":"+cfg.App.Router.GQL.Port, handler))
}
