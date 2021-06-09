package slang

import (
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"math"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const slangCmd = "slang"

func Lint(filename string) (diagnostics []protocol.Diagnostic, cmdText string, err error) {
	dir := filepath.Dir(filename)
	base := filepath.Base(filename)
	cmd := exec.Command(slangCmd, "--lint-only", "--quiet", "--color-diagnostics=false", base)
	cmd.Dir = dir
	cmdText = cmd.String()
	data, _ := cmd.CombinedOutput()
	lines := util.SplitLines(util.DecodeUTF16(data))
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
		} else if strings.Contains(message, "note") {
			severity = protocol.DiagnosticSeverityInformation
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
			Source:             slangCmd,
			Message:            message,
			Tags:               nil,
			RelatedInformation: nil,
			Data:               nil,
		})
	}
	return
}
