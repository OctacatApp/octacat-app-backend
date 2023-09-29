package main

import (
	"fmt"

	"github.com/irdaislakhuafa/octacat-app-backend/src/business/connection"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/domain"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/handler/gql"
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
	psqlDB := connection.NewPostgreSQL(*cfg)

	// init domain
	domain := domain.New(psqlDB)

	// init usecase
	usecase := usecase.New(*cfg, domain)

	// init and run graphql server
	gql.InitAndRun(*cfg, usecase)

	fmt.Printf("cfg: %v\n", *cfg)
}
