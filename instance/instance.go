package instance

import (
	"errors"
	"fmt"
	"github.com/bukka/fpmt/client"
	"github.com/bukka/fpmt/server"
	"github.com/hashicorp/packer/common/json"
	"io/ioutil"
)

type ClientConfig struct {
	Host   string
	Port   string
	Script string
}

type ServerConfig struct {
	Executable string
	ConfigFile string
}

type Config struct {
	Client ClientConfig
	Server ServerConfig
}

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

func (i *Instance) Run() error {
	var config Config
	jsonSpec, err := ioutil.ReadFile(i.Spec)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid spec file %s", i.Spec))
	}

	if err := json.Unmarshal(jsonSpec, &config); err != nil {
		return errors.New(fmt.Sprintf("Invalid JSON in spec file: %s", err.Error()))
	}

	c := transformClient(config.Client)
	s := transformServer(config.Server)

	if err := s.Run("start"); err != nil {
		return err
	}
	if err := c.Run("get"); err != nil {
		return err
	}

	return nil
}
