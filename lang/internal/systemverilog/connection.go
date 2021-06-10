package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"path/filepath"
)

// Initialize implements initialize method.
// https://microsoft.github.io/language-server-protocol/specification#initialize
func (o *Handler) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
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
func (o *Handler) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	o.ShowMessage(fmt.Sprintf("%s started.", ServerName))
	o.workspacePath = "."
	workspaceFolders, err := o.Client.WorkspaceFolders(ctx)
	if err != nil {
		return err
	}
	if workspaceFolders == nil || len(workspaceFolders) == 0 {
		o.ShowWarning("No workspace folder found. Using default settings.")
	} else {
		configFound := false
		for _, ws := range workspaceFolders {
			o.workspacePath = protocol.URI(ws.URI).Filename()
			configFilePath := filepath.Join(o.workspacePath, ConfigFilename)
			if util.Exists(configFilePath) {
				configFound = true
				o.ShowInfo(fmt.Sprintf("Using config file '%s'", configFilePath))
				break
			}
		}
		if !configFound {
			o.ShowWarning("No config file found. Using default settings.")
		}
	}
	return nil
}

// Shutdown implements shutdown method.
// https://microsoft.github.io/language-server-protocol/specification#shutdown
func (o *Handler) Shutdown(ctx context.Context) (err error) {
	o.ShowMessage(fmt.Sprintf("%s shutdown.", ServerName))
	return nil
}

// Exit implements exit method.
// https://microsoft.github.io/language-server-protocol/specification#exit
func (o *Handler) Exit(ctx context.Context) (err error) {
	o.ShowMessage(fmt.Sprintf("%s exited.", ServerName))
	return nil
}
