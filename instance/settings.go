package instance

import "fmt"

// Connection represents a single connection
type Connection struct {
	Host string
	Port string
	Path string
}

// Request represent a request for connection
type Request struct {
	Connection *Connection
	Script     string
}

// Server defines a server
type Server struct {
	Executable     string
	ConfigFile     string
	ConfigTemplate string
}

// Settings contains all settings for the actions
type Settings struct {
	Connections map[string]Connection
	Requests    map[string]Request
	Server      Server
}

// CreateSettings creates a new settings from the config.
func CreateSettings(sc *SettingsConfig) (*Settings, error) {
	if sc == nil {
		return nil, fmt.Errorf("SettingsConfig is nil")
	}

	server := Server{
		ConfigTemplate: "php-fpm.tmpl",
		ConfigFile:     "php-fpm.conf",
		Executable:     "/usr/local/sbin/php-fpm",
	}
	if srvc := sc.Server; srvc != nil {
		if srvc.ConfigFile != "" {
			server.ConfigFile = srvc.ConfigFile
		}
		if srvc.ConfigTemplate != "" {
			server.ConfigTemplate = srvc.ConfigTemplate
		}
		if srvc.Executable != "" {
			server.Executable = srvc.Executable
		}
	}

	settings := Settings{
		Connections: map[string]Connection{},
		Requests:    map[string]Request{},
		Server:      server,
	}

	return &settings, nil
}
