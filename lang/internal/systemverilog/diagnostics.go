package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
)

// DidOpen implements textDocument/didOpen method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didOpen
func (o *Handler) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	go o.publishDiagnostics(params.TextDocument.URI)
	return nil
}

// DidSave implements textDocument/didSave method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didSave
func (o *Handler) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	go o.publishDiagnostics(params.TextDocument.URI)
	return nil
}

func (o *Handler) publishDiagnostics(uri protocol.DocumentURI) {
	diagnostics, cmd, err := verible.Lint(uri.Filename())
	if err != nil {
		o.LogError(fmt.Sprintf("Failed to lint file '%s', error '%e'", uri.Filename(), err))
	}
	o.LogMessage(cmd)

	err = o.Client.PublishDiagnostics(context.TODO(), &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Version:     0,
		Diagnostics: diagnostics,
	})

	if err != nil {
		o.LogError(fmt.Sprintf("Failed to publish diagnostics for file '%s', error '%e'", uri.Filename(), err))
	}
}
