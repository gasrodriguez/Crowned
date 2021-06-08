package verible

import (
	"os/exec"
)

const formatCmd = "verible-verilog-format"

func Format(uri string) (output string, err error) {
	bytes, err := exec.Command(formatCmd, uri).Output()
	return string(bytes), err
}
