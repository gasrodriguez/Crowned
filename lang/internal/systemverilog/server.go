package systemverilog

import (
	"github.com/gasrodriguez/crowned/internal/util"
	"github.com/gasrodriguez/crowned/pkg/server"
	"go.lsp.dev/protocol"
)

type Server struct {
	server.Server
}

func (s *Server) log(message string) {
	err := s.Client.LogMessage(s.Ctx, &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeLog,
	})
	util.CheckError(err)
}

func (s *Server) error(message string) {
	err := s.Client.LogMessage(s.Ctx, &protocol.LogMessageParams{
		Message: message,
		Type:    protocol.MessageTypeError,
	})
	util.CheckError(err)
}

func NewServer() *Server {
	return &Server{}
}
