package instance

import (
	"reflect"
	"testing"
)

func TestCreateSettings(t *testing.T) {
	tables := []struct {
		c *SettingsConfig
		s *Settings
		e string
	}{
		{nil, nil, "SettingsConfig is nil"},
		{
			&SettingsConfig{},
			&Settings{
				Connections: map[string]Connection{},
				Requests:    map[string]Request{},
				Server: Server{
					ConfigFile:     "php-fpm.conf",
					ConfigTemplate: "php-fpm.tmpl",
					Executable:     "/usr/local/sbin/php-fpm",
				},
			},
			"",
		},
		{
			&SettingsConfig{
				Server: &ServerConfig{
					ConfigFile:     "fpm.conf",
					ConfigTemplate: "fpm.tmpl",
				},
			},
			&Settings{
				Connections: map[string]Connection{},
				Requests:    map[string]Request{},
				Server: Server{
					ConfigFile:     "fpm.conf",
					ConfigTemplate: "fpm.tmpl",
					Executable:     "/usr/local/sbin/php-fpm",
				},
			},
			"",
		},
		{
			&SettingsConfig{
				Server: &ServerConfig{
					Executable: "/usr/local/sbin/php-fpmi",
				},
			},
			&Settings{
				Connections: map[string]Connection{},
				Requests:    map[string]Request{},
				Server: Server{
					ConfigFile:     "php-fpm.conf",
					ConfigTemplate: "php-fpm.tmpl",
					Executable:     "/usr/local/sbin/php-fpmi",
				},
			},
			"",
		},
		{
			&SettingsConfig{
				Connection: &ConnectionConfig{
					Port: "9081",
				},
				Request: &RequestConfig{
					Script: "test.php",
				},
			},
			&Settings{
				Connections: map[string]Connection{
					"default": Connection{
						Host: "127.0.0.1",
						Port: "9081",
					},
				},
				Requests: map[string]Request{
					"default": Request{
						Connection: &Connection{
							Host: "127.0.0.1",
							Port: "9081",
						},
						Script: "test.php",
					},
				},
				Server: Server{
					ConfigFile:     "php-fpm.conf",
					ConfigTemplate: "php-fpm.tmpl",
					Executable:     "/usr/local/sbin/php-fpm",
				},
			},
			"",
		},
		{
			&SettingsConfig{
				Connections: &map[string]ConnectionConfig{
					"conn": ConnectionConfig{
						Host: "192.168.1.137",
						Port: "9091",
					},
				},
				Request: &RequestConfig{
					Connection: "conn",
					Script:     "testx.php",
				},
			},
			&Settings{
				Connections: map[string]Connection{
					"conn": Connection{
						Host: "192.168.1.137",
						Port: "9091",
					},
				},
				Requests: map[string]Request{
					"default": Request{
						Connection: &Connection{
							Host: "192.168.1.137",
							Port: "9091",
						},
						Script: "testx.php",
					},
				},
				Server: Server{
					ConfigFile:     "php-fpm.conf",
					ConfigTemplate: "php-fpm.tmpl",
					Executable:     "/usr/local/sbin/php-fpm",
				},
			},
			"",
		},
	}

	for _, table := range tables {
		s, e := CreateSettings(table.c)
		if e != nil {
			if table.e == "" {
				t.Errorf("Expected no error but instead error '%s' returned",
					e.Error())
			} else if e.Error() != table.e {
				t.Errorf("Expected error '%s' but instead error '%s' returned",
					table.e, e.Error())
			}
			continue
		}
		if table.e != "" {
			t.Errorf("Expected error '%s' but no error returned", table.e)
		}
		if !reflect.DeepEqual(s, table.s) {
			t.Errorf("The settings %s does not match expected %s", s, table.s)
		}
	}
}
