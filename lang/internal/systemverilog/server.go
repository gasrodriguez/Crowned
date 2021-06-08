package systemverilog

import (
	"context"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
)

type Server struct {
	protocol.Server
	Client protocol.Client
	Ctx    context.Context
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
