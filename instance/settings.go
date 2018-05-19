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
	return nil, fmt.Errorf("SettingsConfig is nil")
}
