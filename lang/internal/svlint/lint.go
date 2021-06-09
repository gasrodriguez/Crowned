package svlint

import (
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"math"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const lintCmd = "svlint"

func Lint(filename string) (diagnostics []protocol.Diagnostic, cmdText string, err error) {
	dir := filepath.Dir(filename)
	base := filepath.Base(filename)
	cmd := exec.Command(lintCmd, "-1", base)
	cmd.Dir = dir
	cmdText = cmd.String()
	data, _ := cmd.CombinedOutput()
	lines := util.SplitLines(data)
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
		terms[2] = strings.Split(terms[2], "\u001B[37m")[0]
		colNum, err := strconv.Atoi(terms[2])
		if err != nil {
			continue
		}
		message := terms[len(terms)-1]
		message = strings.Replace(message, "\u001B[0m", "", 1)
		message = strings.TrimSpace(message)
		severity := protocol.DiagnosticSeverityHint

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
