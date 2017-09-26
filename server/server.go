package server

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Server struct {
	FpmBinary string
	FpmConfig string
}

func (s *Server) start() error {
	var fpmBinary string
	configPath, _ := filepath.Abs(s.FpmConfig)
	if s.FpmBinary[0] != '/' {
		fpmBinary = "/usr/local/sbin/" + s.FpmBinary
	} else {
		fpmBinary = s.FpmBinary
	}
	cmd := exec.Command(fpmBinary, "-F", "-y", configPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *Server) Run(action string) error {
	switch action {
	case "start":
		return s.start()
	default:
		return errors.New(fmt.Sprintf("Unknown action %s", action))
	}
}
