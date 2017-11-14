package server

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
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
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("Sending SIGTERM to process %d\n", cmd.Process.Pid)
		cmd.Process.Signal(syscall.SIGTERM)
	}()
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
