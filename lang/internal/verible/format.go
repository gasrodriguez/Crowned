package verible

import (
	"os/exec"
	"path/filepath"
)

const formatCmd = "verible-verilog-format"

func Format(cwd, filename string, args []string) (output string, cmdText string, err error) {
	relPath, err := filepath.Rel(cwd, filename)
	if err != nil {
		relPath = filename
	}
	args = append(args, relPath)
	cmd := exec.Command(formatCmd, args...)
	cmd.Dir = cwd
	cmdText = cmd.String()
	data, err := cmd.Output()
	return string(data), cmdText, err
}
