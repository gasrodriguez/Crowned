package main

import (
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"os"
)

func main() {
	svHandler := systemverilog.NewHandler()
	svHandler.Run(svHandler, os.Args)
}
