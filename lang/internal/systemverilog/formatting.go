package systemverilog

import (
	"context"
	"github.com/gasrodriguez/crowned/internal/util"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
	"math"
)

func (s *Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	endLine, err := util.LineCounter(params.TextDocument.URI.Filename())
	if err != nil {
		return
	}
	newText, err := verible.Format(params.TextDocument.URI.Filename())
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
