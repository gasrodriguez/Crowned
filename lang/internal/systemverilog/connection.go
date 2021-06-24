package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/util"
	"go.lsp.dev/protocol"
	"path/filepath"
)

func (o *Handler) doInitialize(_ context.Context, _ *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
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

func (o *Handler) doInitialized(ctx context.Context, _ *protocol.InitializedParams) (err error) {
	o.ShowMessage(fmt.Sprintf("%s started.", ServerName))
	o.workspacePath = "."
	workspaceFolders, err := o.Client.WorkspaceFolders(ctx)
	if err != nil {
		return err
	}
	if workspaceFolders == nil || len(workspaceFolders) == 0 {
		o.ShowWarning("No workspace folder found.\nUsing default settings.")
	} else {
		configFound := false
		for _, ws := range workspaceFolders {
			o.workspacePath = protocol.URI(ws.URI).Filename()
			configFilePath := filepath.Join(o.workspacePath, ConfigFilename)
			if util.Exists(configFilePath) {
				configFound = true
				break
			}
		}
		if !configFound {
			o.ShowWarning("No config file found.\nUsing default settings.")
		}
	}
	config := o.loadConfig()
	o.files.ScanWorkspace(o.workspacePath, config.General.Filters)
	return nil
}
