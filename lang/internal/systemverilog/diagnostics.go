package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/slang"
	"github.com/gasrodriguez/crowned/internal/svlint"
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
	go o.publishDiagnostics(context.TODO(), params.TextDocument.URI)
	return nil
}

func (o *Handler) publishDiagnostics(ctx context.Context, uri protocol.DocumentURI) {
	diagnostics := make([]protocol.Diagnostic, 0)
	config := o.loadConfig()

	if config.Slang.Enabled {
		diagnosticsSlang, cmd, err := slang.Lint(o.workspacePath, uri.Filename(), config.General.Includes, config.Slang.Arguments)
		o.LogMessage(cmd)
		if err != nil {
			o.LogError(fmt.Sprintf("Failed to lint file '%s', error '%s'", uri.Filename(), err.Error()))
		} else {
			diagnostics = append(diagnostics, diagnosticsSlang...)
		}
	}

	if config.Svlint.Enabled {
		diagnosticsSvlint, cmd, err := svlint.Lint(o.workspacePath, uri.Filename(), config.General.Includes, config.Svlint.Arguments)
		o.LogMessage(cmd)
		if err != nil {
			o.LogError(fmt.Sprintf("Failed to lint file '%s', error '%s'", uri.Filename(), err.Error()))
		} else {
			diagnostics = append(diagnostics, diagnosticsSvlint...)
		}
	}

	if config.Verible.Lint.Enabled {
		diagnosticsVerible, cmd, err := verible.Lint(o.workspacePath, uri.Filename(), config.Verible.Lint.Arguments, config.Verible.Lint.Rules)
		o.LogMessage(cmd)
		if err != nil {
			o.LogError(fmt.Sprintf("Failed to lint file '%s', error '%s'", uri.Filename(), err.Error()))
		} else {
			diagnostics = append(diagnostics, diagnosticsVerible...)
		}
	}

	err := o.Client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Version:     0,
		Diagnostics: diagnostics,
	})

	if err != nil {
		o.LogError(fmt.Sprintf("Failed to publish diagnostics for file '%s', error '%e'", uri.Filename(), err))
	}
}

// DidClose implements textDocument/didClose method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didClose
func (o *Handler) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	uri := params.TextDocument.URI
	err = o.Client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Version:     0,
		Diagnostics: []protocol.Diagnostic{},
	})

	if err != nil {
		o.LogError(fmt.Sprintf("Failed to clear diagnostics for file '%s', error '%e'", uri.Filename(), err))
	}
	return
}
