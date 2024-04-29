package configreader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

func ReadConfigFile(filePath string) (*config.AppConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}

	cfgBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInternalServerError, err.Error())
	}

	cfg := new(config.AppConfig)
	if err := json.Unmarshal(cfgBytes, cfg); err != nil {
		return nil, errors.NewWithCode(codes.CodeJSONUnmarshalError, err.Error())
	}

	return cfg, nil
}
