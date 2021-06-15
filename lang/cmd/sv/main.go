package main

import (
	"flag"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/systemverilog"
	"os"
	"strings"
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
	return strings.Join([]string{
		fmt.Sprintf("%s (%s)", systemverilog.ServerName, systemverilog.ServerVersion),
		"Git branch: " + gitBranch,
		"Git Commit: " + gitCommit,
		"Build time: " + buildTime,
	}, eol)
}
