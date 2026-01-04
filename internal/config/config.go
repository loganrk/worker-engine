package config

import (
	"fmt"

	"github.com/loganrk/worker-engine/internal/core/port"
	"github.com/spf13/viper"
)

type File struct {
	Name string
	Ext  string
}

func StartConfig(path string, file File) (port.ConfApp, error) {
	var appConfig app

	var viperIns = viper.New()

	viperIns.AddConfigPath(path)
	viperIns.SetConfigName(file.Name)
	viperIns.AddConfigPath(".")
	viperIns.SetConfigType(file.Ext)

	if err := viperIns.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err := viperIns.Unmarshal(&appConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return appConfig, nil
}

func (a app) GetAppName() string {
	return a.Application.Name
}

func (a app) GetLogger() port.ConfLogger {
	return a.Logger
}

func (a app) GetUser() port.ConfUser {
	return a.User
}

func (a app) GetKafka() port.ConfKafka {
	return a.Kafka
}
func (a app) GetEmail() port.ConfEmail {
	return a.Email
}
