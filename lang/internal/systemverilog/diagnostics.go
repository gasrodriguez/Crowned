package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/slang"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
)

// DidOpen implements textDocument/didOpen method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didOpen
func (o *Handler) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	go o.publishDiagnostics(ctx, params.TextDocument.URI)
	return nil
}

// DidSave implements textDocument/didSave method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didSave
func (o *Handler) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	go o.publishDiagnostics(ctx, params.TextDocument.URI)
	return nil
}

func (o *Handler) publishDiagnostics(ctx context.Context, uri protocol.DocumentURI) {
	diagnosticsVerible, cmd, err := verible.Lint(uri.Filename())
	if err != nil {
		o.LogError(fmt.Sprintf("Failed to lint file '%s', error '%s'", uri.Filename(), err.Error()))
	}
	o.LogMessage(cmd)

	diagnosticsSlang, cmd, err := slang.Lint(uri.Filename())
	if err != nil {
		o.LogError(fmt.Sprintf("Failed to lint file '%s', error '%s'", uri.Filename(), err.Error()))
	}
	o.LogMessage(cmd)

	err = o.Client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Version:     0,
		Diagnostics: append(diagnosticsVerible, diagnosticsSlang...),
	})

	if err != nil {
		o.LogError(fmt.Sprintf("Failed to publish diagnostics for file '%s', error '%e'", uri.Filename(), err))
	}
}
