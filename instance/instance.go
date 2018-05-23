package instance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConnectionConfig sets connection parameters
type ConnectionConfig struct {
	Host string
	Port string
	Path string
}

// RequestConfig is for configuring request parameters.
type RequestConfig struct {
	Connection string
	Script     string
}

// ServerConfig is for configuring server parameters.
type ServerConfig struct {
	Executable     string
	ConfigFile     string
	ConfigTemplate string
}

// SettingsConfig wraps all settings
type SettingsConfig struct {
	Connection  *ConnectionConfig
	Connections *map[string]ConnectionConfig
	Request     *RequestConfig
	Requests    *map[string]RequestConfig
	Server      *ServerConfig
}

// Config is the main section.
type Config struct {
	Settings SettingsConfig
	Actions  []interface{}
}

// Instance is for input parameters.
type Instance struct {
	Spec string
}

// Run the instance.
func (i *Instance) Run() error {
	var config Config
	var settings *Settings
	jsonSpec, err := ioutil.ReadFile(i.Spec)
	if err != nil {
		return fmt.Errorf("Invalid spec file %s", i.Spec)
	}

	if err := json.Unmarshal(jsonSpec, &config); err != nil {
		return fmt.Errorf("Invalid JSON in spec file: %s", err.Error())
	}
	if settings, err = CreateSettings(&config.Settings); err != nil {
		return err
	}

	for _, actionConfig := range config.Actions {
		action, err := CreateAction(actionConfig)
		if err != nil {
			return err
		}
		if err := action.Run(settings); err != nil {
			return err
		}
	}

	return nil
}
