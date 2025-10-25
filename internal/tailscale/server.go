package tailscale

import (
	"context"

	"tailscale.com/tsnet"
)

type Server struct {
	TSServer *tsnet.Server
}

func NewServer(hostname string, authKey string, controlURL string) *Server {
	s := &tsnet.Server{
		Hostname:   hostname,
		Dir:        "/var/lib/tsnet",
		AuthKey:    authKey,
		ControlURL: controlURL,
	}
	return &Server{TSServer: s}
}

func (server *Server) Connect() error {
	_, err := server.TSServer.Up(context.Background())
	return err
}

func (server *Server) Close() error {
	return server.TSServer.Close()
}
