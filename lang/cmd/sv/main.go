package main

import (
	"github.com/gasrodriguez/crowned/internal/svlint"
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"os"
)

func main() {
	svlint.Lint("D:\\Proyectos\\Plugins\\crowned\\demo", "D:\\Proyectos\\Plugins\\crowned\\demo\\tb_build\\rtl\\spi\\rtl\\verilog\\spi_top.v", []string{"tb_build/rtl/spi/rtl/verilog"}, nil)
	svHandler := systemverilog.NewHandler()
	svHandler.Run(svHandler, os.Args)
}
