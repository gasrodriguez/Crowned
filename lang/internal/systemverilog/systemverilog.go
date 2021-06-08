package systemverilog

import (
	"github.com/gasrodriguez/crowned/pkg/handler"
)

const (
	ServerName     = "Crowned SystemVerilog Language Server"
	ServerVersion  = "0.0.1"
	ConfigFilename = "crowned.toml"
)

type SystemVerilog struct {
	handler.Handler
}

func NewServer() *SystemVerilog {
	return &SystemVerilog{}
}
