package server

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Server struct {
	FpmBinary string
	FpmConfig string
}

func (s *Server) start() error {
	cmd := exec.Command(s.FpmBinary, "-F", "-y", s.FpmConfig)
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
