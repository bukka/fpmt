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
