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

const (
	lintCmd = "svlint"

	errorPrefix  = "Error: "
	failPrefix   = "Fail: "
	hintPrefix   = "hint  : "
	reasonPrefix = "reason: "
)

func Lint(cwd, filename string, includes []string, args []string) (diagnostics []protocol.Diagnostic, cmdText string, err error) {
	relPath, err := filepath.Rel(cwd, filename)
	if err != nil {
		relPath = filename
	}
	for _, include := range includes {
		args = append(args, "--include="+include)
	}
	args = append(args, relPath)
	cmd := exec.Command(lintCmd, args...)
	cmd.Dir = cwd
	cmdText = cmd.String()
	data, err := cmd.CombinedOutput()
	lines := util.SplitLines(util.Decolorize(data))
	diagnostics = make([]protocol.Diagnostic, 0)

	index := 0
	code := ""
	lineNum := 1
	colNum := 1
	message := ""
	severity := protocol.DiagnosticSeverityHint

	for _, line := range lines {
		switch index {
		case 0:
			lineNum = 1
			colNum = 1
			message = ""
			if strings.HasPrefix(line, errorPrefix) {
				code = strings.TrimPrefix(line, errorPrefix)
				message = code
				severity = protocol.DiagnosticSeverityError
				index = 1
			} else if strings.HasPrefix(line, failPrefix) {
				code = strings.TrimPrefix(line, failPrefix)
				severity = protocol.DiagnosticSeverityHint
				index = 1
			}
			continue

		case 1:
			terms := strings.Split(line, ":")
			if len(terms) < 4 || terms[0] != relPath {
				index = 0
				continue
			}
			lineNum, err = strconv.Atoi(terms[1])
			if err != nil {
				continue
			}
			colNum, err = strconv.Atoi(terms[2])
			if err != nil {
				continue
			}
			if severity == protocol.DiagnosticSeverityError {
				index = 0
				break
			} else {
				index++
				continue
			}

		case 4:
			message = hintPrefix + util.StringListLast(strings.Split(line, hintPrefix)) + "\n"
			index++
			continue

		case 5:
			message += reasonPrefix + util.StringListLast(strings.Split(line, reasonPrefix)) + "\n"
			index = 0
			break

		default:
			index++
			continue
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
			Code:               code,
			Source:             lintCmd,
			Message:            message,
			Tags:               nil,
			RelatedInformation: nil,
			Data:               nil,
		})
	}
	return
}
