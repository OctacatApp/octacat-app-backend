package main

import (
	"fmt"

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

	fmt.Printf("cfg: %v\n", *cfg)
}
