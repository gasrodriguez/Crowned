package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
)

// DidOpen implements textDocument/didOpen method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didOpen
func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	go s.publishDiagnostics(params.TextDocument.URI)
	return nil
}

// DidSave implements textDocument/didSave method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didSave
func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	go s.publishDiagnostics(params.TextDocument.URI)
	return nil
}

func (s *Server) publishDiagnostics(uri protocol.DocumentURI) {
	diagnostics, cmd, err := verible.Lint(uri.Filename())
	if err != nil {
		s.LogError(fmt.Sprintf("Failed to lint file '%s', error '%e'", uri.Filename(), err))
	}
	s.LogMessage(cmd)

	err = s.Client.PublishDiagnostics(context.TODO(), &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Version:     0,
		Diagnostics: diagnostics,
	})

	if err != nil {
		s.LogError(fmt.Sprintf("Failed to publish diagnostics for file '%s', error '%e'", uri.Filename(), err))
	}
}
