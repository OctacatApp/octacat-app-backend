package main

import (
	"fmt"

	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/files"
)

const (
	configPath = "./etc/cfg"
	configFile = "conf.json"
)

func main() {
	if !files.IsFileExist(configPath) {
		panic(fmt.Sprintf("dir %v doesn't exist!", configPath))
	}
}
