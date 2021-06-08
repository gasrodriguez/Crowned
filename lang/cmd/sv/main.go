package main

import (
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"os"
)

func main() {
	svServer := systemverilog.NewServer()
	svServer.Run(svServer, os.Args)
}
