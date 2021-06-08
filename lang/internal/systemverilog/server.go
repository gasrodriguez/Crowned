package systemverilog

import (
	"github.com/gasrodriguez/crowned/pkg/server"
)

type Server struct {
	server.Server
}

func NewServer() *Server {
	return &Server{}
}
