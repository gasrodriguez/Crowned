package main

import (
	"fmt"
	"github.com/gasrodriguez/crowned/internal/systemverilog"
)

func main() {
	files := systemverilog.NewFiles()
	files.ScanWorkspace("../demo", []string{})

	args := "--include-directory=C:\\FpgaTools\\Mentor\\questasim64_10.5c\\verilog_src\\uvm-1.2/src --libdir=C:\\FpgaTools\\Mentor\\questasim64_10.5c\\verilog_src\\uvm-1.2/src -Wredef-macro"
	for _, inc := range files.CompileFiles() {
		args += " -I=" + inc
	}
	for _, comp := range files.IncludeDirs() {
		args += " " + comp
	}
	fmt.Println(args)
}
