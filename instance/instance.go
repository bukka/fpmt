package instance

import (
	"github.com/bukka/fpmt/client"
	"github.com/bukka/fpmt/server"
)

type Config struct {
	client client.Client
	server server.Server
}
