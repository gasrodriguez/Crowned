package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"path"
)

// Initialize implements initialize method.
// https://microsoft.github.io/language-server-protocol/specification#initialize
func (o *SystemVerilog) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncOptions{
				OpenClose:         true,
				Change:            0,
				WillSave:          false,
				WillSaveWaitUntil: false,
				Save: &protocol.SaveOptions{
					IncludeText: false,
				},
			},
			HoverProvider:                    nil,
			CompletionProvider:               nil,
			SignatureHelpProvider:            nil,
			DeclarationProvider:              nil,
			DefinitionProvider:               nil,
			TypeDefinitionProvider:           nil,
			ImplementationProvider:           nil,
			ReferencesProvider:               nil,
			DocumentHighlightProvider:        nil,
			DocumentSymbolProvider:           nil,
			WorkspaceSymbolProvider:          nil,
			CodeActionProvider:               nil,
			CodeLensProvider:                 nil,
			DocumentFormattingProvider:       true,
			DocumentRangeFormattingProvider:  nil,
			DocumentOnTypeFormattingProvider: nil,
			RenameProvider:                   nil,
			DocumentLinkProvider:             nil,
			ColorProvider:                    nil,
			FoldingRangeProvider:             nil,
			SelectionRangeProvider:           nil,
			ExecuteCommandProvider:           nil,
			Workspace:                        nil,
			LinkedEditingRangeProvider:       nil,
			CallHierarchyProvider:            nil,
			SemanticTokensProvider:           nil,
			MonikerProvider:                  nil,
			Experimental:                     nil,
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    ServerName,
			Version: ServerVersion,
		},
	}, nil
}

// Initialized implements initialized method.
// https://microsoft.github.io/language-server-protocol/specification#initialized
func (o *SystemVerilog) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	o.ShowInfo(fmt.Sprintf("%s started.", ServerName))
	workspaceFolders, err := o.Client.WorkspaceFolders(ctx)
	if err != nil {
		return err
	}
	if workspaceFolders == nil || len(workspaceFolders) == 0 {
		o.ShowWarning("No workspace folder found. Starting in single-file mode.")
	} else {
		var configFile *string
		for _, ws := range workspaceFolders {
			cfg := path.Join(protocol.URI(ws.URI).Filename(), ConfigFilename)
			if util.Exists(cfg) {
				configFile = &cfg
			}
		}
		if configFile != nil {
			o.ShowWarning("No config file found. Starting in single-file mode.")
		} else {
			// ToDo: config file not supported yet :)
			o.ShowWarning("Config file not yet supported. Starting in single-file mode.")
		}
	}
	return nil
}

// Shutdown implements shutdown method.
// https://microsoft.github.io/language-server-protocol/specification#shutdown
func (o *SystemVerilog) Shutdown(ctx context.Context) (err error) {
	o.ShowInfo(fmt.Sprintf("%s shutdown.", ServerName))
	return nil
}

// Exit implements exit method.
// https://microsoft.github.io/language-server-protocol/specification#exit
func (o *SystemVerilog) Exit(ctx context.Context) (err error) {
	o.ShowInfo(fmt.Sprintf("%s exited.", ServerName))
	return nil
}
