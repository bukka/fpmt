package server

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"text/template"
)

type Server struct {
	FpmBinary         string
	FpmConfig         string
	FpmConfigTemplate string
}

// Get config path.
func (s *Server) getConfigPath() (string, error) {
	// get absolute config path
	configPath, err := filepath.Abs(s.FpmConfig)
	if err != nil {
		return "", err
	}
	// if there is no template, return config path
	if len(s.FpmConfigTemplate) == 0 {
		return configPath, nil
	}
	configTemplatePath, err := filepath.Abs(s.FpmConfigTemplate)
	if err != nil {
		return "", err
	}
	// create template instance from the supplied template path
	configTemplate, err := template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"nrange": func(n int) (stream chan int) {
			stream = make(chan int)
			go func() {
				for i := 1; i <= n; i++ {
					stream <- i
				}
				close(stream)
			}()
			return
		},
	}).ParseFiles(configTemplatePath)
	if err != nil {
		return "", err
	}
	// create io.Writer from the config file path
	configFile, err := os.Create(configPath)
	if err != nil {
		return "", err
	}
	defer configFile.Close()
	// execute the template and save the result to config file
	if err := configTemplate.ExecuteTemplate(configFile, filepath.Base(configTemplatePath), s); err != nil {
		return "", err
	}
	// return the config file
	return configPath, nil
}

// Start the server.
func (s *Server) start() error {
	// get config path
	configPath, err := s.getConfigPath()
	if err != nil {
		return err
	}
	// get fpm binary to execute
	var fpmBinary string
	if s.FpmBinary[0] != '/' {
		fpmBinary = "/usr/local/sbin/" + s.FpmBinary
	} else {
		fpmBinary = s.FpmBinary
	}
	// execute the fpm binary in foreground with the config file
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

// Run the server action.
func (s *Server) Run(action string) error {
	switch action {
	case "start":
		return s.start()
	default:
		return fmt.Errorf("unknown action %s", action)
	}
}
