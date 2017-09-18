package server

import (
	"errors"
	"fmt"
	"os/exec"
)

type Server struct {
	FpmBinary string
	FpmConfig string
}

func (s *Server) start() error {
	cmd := exec.Command(s.FpmBinary, "-F", "-y", s.FpmConfig)
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
