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

	server, err := createServer(sc)
	if err != nil {
		return nil, err
	}
	conns, err := createConnections(sc)
	if err != nil {
		return nil, err
	}
	reqs, err := createRequests(sc, conns)
	if err != nil {
		return nil, err
	}

	settings := Settings{
		Connections: conns,
		Requests:    reqs,
		Server:      server,
	}

	return &settings, nil
}

func createConnections(sc *SettingsConfig) (map[string]Connection, error) {
	if sc.Connection == nil {
		return map[string]Connection{}, nil
	}
	conn := Connection{
		Host: "127.0.0.1",
	}
	if sc.Connection.Port != "" {
		conn.Port = sc.Connection.Port
	}
	return map[string]Connection{"default": conn}, nil
}

func createRequests(sc *SettingsConfig, conns map[string]Connection) (map[string]Request, error) {
	if sc.Request == nil {
		return map[string]Request{}, nil
	}
	conn := conns["default"]
	req := Request{
		Connection: &conn,
	}
	if sc.Request.Script != "" {
		req.Script = sc.Request.Script
	}
	return map[string]Request{"default": req}, nil
}

func createServer(sc *SettingsConfig) (Server, error) {
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

	return server, nil
}
