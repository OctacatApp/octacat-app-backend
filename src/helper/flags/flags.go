package flags

import (
	"flag"
	"fmt"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/files"
)

// default env is "local"
func ParseFlags(configPath, configFile string) (*string, error) {
	env := ""
	flag.StringVar(&env, "env", "local", fmt.Sprintf("mode from dir name in %v/", configPath))
	flag.Parse()

	if dir := fmt.Sprintf("%v/%v", configPath, env); !files.IsFileExist(dir) {
		return nil, errors.NewWithCode(codes.CodeNotFound, "env '%v' not found in dir '%v/' \n", env, configPath)
	} else {
		if cfgFile := fmt.Sprintf("%v/%v", dir, configFile); !files.IsFileExist(cfgFile) {
			return nil, errors.NewWithCode(codes.CodeNotFound, "'%v' not found in '%v/'\n", configFile, dir)
		}
	}
	return &env, nil
}
