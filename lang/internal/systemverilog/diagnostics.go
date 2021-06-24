package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/slang"
	"github.com/gasrodriguez/crowned/internal/svlint"
	"github.com/gasrodriguez/crowned/internal/verible"
	"go.lsp.dev/protocol"
)

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
		diagnosticsVerible, cmd, err := verible.Lint(o.workspacePath, uri.Filename(), config.Verible.Lint.Arguments)
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

func (o *Handler) clearDiagnostics(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
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
