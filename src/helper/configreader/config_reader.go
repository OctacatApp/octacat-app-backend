package configreader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

func ReadConfigFile(filePath string) (*config.AppConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	cfgBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := new(config.AppConfig)
	if err := json.Unmarshal(cfgBytes, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
