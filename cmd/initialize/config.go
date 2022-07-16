package initialize

import (
	"context"

	"github.com/pikomonde/i-view-nityo/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configJSONFileName = "config.yaml"
)

// NewConfig Read and process config file
func NewConfig(ctx context.Context, logger *log.Entry) (config model.Config, err error) {
	// read from config file
	viper.SetConfigFile(configJSONFileName)
	err = viper.ReadInConfig()
	if err != nil {
		logger.WithField("err", err).Errorln("can't find config file")
		return config, err
	}
	logger.Infoln("Config loaded from config file")

	// unmarshal to the struct
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err).Errorln("error unmarshal config to struct")
		return config, err
	}
	logger.Infoln("Unmarshal config file")

	// setup logging
	switch config.App.Env {
	case model.Env_Local, model.Env_Staging:
		log.SetLevel(log.DebugLevel)
	case model.Env_Production:
		log.SetLevel(log.InfoLevel)
	}

	return
}
