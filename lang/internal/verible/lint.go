package verible

import (
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"math"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const lintCmd = "verible-verilog-lint"

func Lint(cwd, filename string, args []string, rules []string) (diagnostics []protocol.Diagnostic, cmdText string, err error) {
	relPath, err := filepath.Rel(cwd, filename)
	if err != nil {
		relPath = filename
	}
	args = append(args, "--lint_fatal=false")
	args = append(args, "--parse_fatal=false")
	if rules != nil && len(rules) > 0 {
		args = append(args, "--rules="+strings.Join(rules, ","))
	}
	args = append(args, relPath)
	cmd := exec.Command(lintCmd, args...)
	cmd.Dir = cwd
	cmdText = cmd.String()
	bytes, err := cmd.Output()
	if err != nil {
		return
	}
	lines := util.SplitLines(bytes)
	diagnostics = make([]protocol.Diagnostic, 0)
	for _, line := range lines {
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
