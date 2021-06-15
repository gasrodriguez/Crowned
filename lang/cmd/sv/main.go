package main

import (
	"flag"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"os"
)

const (
	eol = "\n"
)

var (
	gitTag    string
	gitBranch string
	gitCommit string
	buildTime string
)

func main() {
	versionFlag := flag.Bool("v", false, "Print the current version and exit")
	flag.Parse()

	systemverilog.ServerVersion = gitTag

	switch {
	case *versionFlag:
		fmt.Println(version())
		return
	}

	svHandler := systemverilog.NewHandler()
	svHandler.Run(svHandler, os.Args)
}

func version() string {
	result := fmt.Sprintf("%s (%s)\n", systemverilog.ServerName, systemverilog.ServerVersion)
	result += "Git branch: " + gitBranch + eol
	result += "Git Commit: " + gitCommit + eol
	result += "Build time: " + buildTime + eol
	return result
}
