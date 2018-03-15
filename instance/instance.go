package instance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bukka/fpmt/client"
	"github.com/bukka/fpmt/server"
)

// ClientConfig is for configuring client parameters.
type ClientConfig struct {
	Host   string
	Port   string
	Script string
}

// ServerConfig is for configuring server parameters.
type ServerConfig struct {
	Executable     string
	ConfigFile     string
	ConfigTemplate string
}

// Config is the main section.
type Config struct {
	Client ClientConfig
	Server ServerConfig
}

// Instance is for input parameters.
type Instance struct {
	Spec string
}

func transformClient(cfg ClientConfig) *client.Client {
	c := &client.Client{
		Host:   cfg.Host,
		Port:   cfg.Port,
		Script: cfg.Script,
	}

	return c
}

func transformServer(cfg ServerConfig) *server.Server {
	s := &server.Server{
		FpmBinary: cfg.Executable,
		FpmConfig: cfg.ConfigFile,
	}

	return s
}

// Run the instance.
func (i *Instance) Run() error {
	var config Config
	jsonSpec, err := ioutil.ReadFile(i.Spec)
	if err != nil {
		return fmt.Errorf("Invalid spec file %s", i.Spec)
	}

	if err := json.Unmarshal(jsonSpec, &config); err != nil {
		return fmt.Errorf("Invalid JSON in spec file: %s", err.Error())
	}

	c := transformClient(config.Client)
	s := transformServer(config.Server)

	if err := s.Run("start"); err != nil {
		return err
	}

	return c.Run("get")
}
