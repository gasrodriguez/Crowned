package systemverilog

import (
	"github.com/gasrodriguez/crowned/pkg/lsp"
)

const (
	ServerName = "Crowned SystemVerilog Language Server"
)

var (
	ServerVersion = "0.0.0"
)

type Handler struct {
	lsp.Handler
	workspacePath string
}

func NewHandler() *Handler {
	return &Handler{}
}
