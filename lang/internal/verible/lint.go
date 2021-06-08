package verible

import (
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"math"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const lintCmd = "verible-verilog-lint"

func Lint(filename string) (diagnostics []protocol.Diagnostic, cmdText string, err error) {
	dir := path.Dir(filename)
	base := path.Base(filename)
	cmd := exec.Command(lintCmd, "--lint_fatal=false", "--parse_fatal=false", base)
	cmd.Dir = dir
	cmdText = cmd.String() + dir + base
	bytes, err := cmd.Output()
	if err != nil {
		return
	}
	cmdText = string(bytes)
	lines := util.SplitLines(bytes)
	diagnostics = make([]protocol.Diagnostic, 0)
	for _, line := range lines {
		// workaround
		line = strings.Replace(line, filename, "", 1)
		terms := strings.Split(line, ":")
		if len(terms) < 4 {
			continue
		}
		lineNum, err := strconv.Atoi(terms[1])
		if err != nil {
			continue
		}
		colNum, err := strconv.Atoi(terms[2])
		if err != nil {
			continue
		}
		message := strings.Join(terms[3:], ":")
		severity := protocol.DiagnosticSeverityWarning
		if strings.Contains(message, "error") {
			severity = protocol.DiagnosticSeverityError
		}

		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      uint32(lineNum - 1),
					Character: uint32(colNum - 1),
				},
				End: protocol.Position{
					Line:      uint32(lineNum - 1),
					Character: math.MaxUint32,
				},
			},
			Severity:           severity,
			Code:               nil,
			CodeDescription:    nil,
			Source:             lintCmd,
			Message:            message,
			Tags:               nil,
			RelatedInformation: nil,
			Data:               nil,
		})
	}
	return
}
