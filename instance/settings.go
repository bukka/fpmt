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
	if sc.Connection == nil && sc.Connections == nil {
		return map[string]Connection{}, nil
	}
	conns := map[string]Connection{}
	if sc.Connections != nil {
		for name, cc := range *sc.Connections {
			if err := createConnection(&cc, conns, name); err != nil {
				return conns, err
			}
		}
	}
	if sc.Connection != nil {
		if err := createConnection(sc.Connection, conns, "default"); err != nil {
			return conns, err
		}
	}

	return conns, nil
}

func createConnection(cc *ConnectionConfig, conns map[string]Connection, name string) error {
	conn := Connection{}
	if cc.Host != "" {
		conn.Host = cc.Host
	}
	if cc.Port != "" {
		conn.Port = cc.Port
		if cc.Host == "" {
			conn.Host = "127.0.0.1"
		}
	}
	if cc.Path != "" {
		if conn.Port != "" {
			return fmt.Errorf("Connection can have either path or host/port not both")
		}
		conn.Path = cc.Path
	}
	conns[name] = conn

	return nil
}

func createRequests(sc *SettingsConfig, conns map[string]Connection) (map[string]Request, error) {
	if sc.Request == nil && sc.Requests == nil {
		return map[string]Request{}, nil
	}
	reqs := map[string]Request{}
	if sc.Requests != nil {
		for name, rc := range *sc.Requests {
			if err := createRequest(&rc, conns, reqs, name); err != nil {
				return reqs, err
			}
		}
	}
	if sc.Request != nil {
		if err := createRequest(sc.Request, conns, reqs, "default"); err != nil {
			return reqs, err
		}
	}

	return reqs, nil
}

func createRequest(rc *RequestConfig, conns map[string]Connection, reqs map[string]Request, name string) error {
	cn := "default"
	if rc.Connection != "" {
		cn = rc.Connection
	}

	conn, ok := conns[cn]
	if !ok {
		return fmt.Errorf("No connection '%s' found in the Request '%s'", cn, name)
	}
	req := Request{
		Connection: &conn,
	}
	if rc.Script != "" {
		req.Script = rc.Script
	}
	reqs[name] = req

	return nil
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
