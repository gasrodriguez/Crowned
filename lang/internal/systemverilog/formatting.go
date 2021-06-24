package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/util"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
	"math"
)

func (o *Handler) doFormatting(_ context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	filename := params.TextDocument.URI.Filename()
	endLine, err := util.LineCounter(filename)
	if err != nil {
		return nil, err
	}
	config := o.loadConfig()
	if !config.Verible.Format.Enabled {
		return
	}
	newText, cmd, err := verible.Format(o.workspacePath, filename, config.Verible.Format.Arguments)
	o.LogMessage(cmd)
	if err != nil {
		o.LogError(fmt.Sprintf("Failed to format file '%s', error '%s'", filename, err.Error()))
		return nil, err
	}

	result = []protocol.TextEdit{
		{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      0,
					Character: 0,
				},
				End: protocol.Position{
					Line:      endLine,
					Character: math.MaxUint32,
				},
			},
			NewText: newText,
		},
	}
	return result, err
}
