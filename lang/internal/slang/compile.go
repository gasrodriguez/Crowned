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

const compileCmd = "slang"

// Compile is similar to lint but works for full project
func Compile(cwd string, compile []string, includes []string, libdirs []string, args []string) (diagnostics []protocol.Diagnostic, cmdText string, err error) {
	args = append(args, "--lint-only")
	args = append(args, "--quiet")
	args = append(args, "--color-diagnostics=false")
	for _, include := range includes {
		args = append(args, "-I="+include)
	}
	for _, libdir := range libdirs {
		args = append(args, "-I="+libdir)
		args = append(args, "--libdir="+libdir)
	}
	for _, filename := range compile {
		relPath, err := filepath.Rel(cwd, filename)
		if err != nil {
			relPath = filename
		}
		args = append(args, relPath)
	}
	cmd := exec.Command(compileCmd, args...)
	cmd.Dir = cwd
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
			Source:             lintCmd,
			Message:            message,
			Tags:               nil,
			RelatedInformation: nil,
			Data:               nil,
		})
	}
	return
}