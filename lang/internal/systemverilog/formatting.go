package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/util"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
	"math"
)

func (o *Handler) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	o.LogMessage("received formatting req")
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

//func (s *Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
//	err = notImplemented("OnTypeFormatting")
//	return
//}
//
//func (s *Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
//	err = notImplemented("RangeFormatting")
//	return
//}
//
