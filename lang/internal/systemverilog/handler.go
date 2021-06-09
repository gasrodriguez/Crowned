package systemverilog

import (
	"github.com/gasrodriguez/crowned/pkg/lsp"
)

const (
	ServerName     = "Crowned SystemVerilog Language Server"
	ServerVersion  = "0.0.1"
	ConfigFilename = "crowned.toml"
)

type Handler struct {
	lsp.Handler
}

func NewHandler() *Handler {
	return &Handler{}
}
