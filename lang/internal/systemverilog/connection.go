package systemverilog

import (
	"context"
	"go.lsp.dev/protocol"
)

// Initialize implements initialize method.
// https://microsoft.github.io/language-server-protocol/specification#initialize
func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
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
			Name:    "SV Language Server",
			Version: "0.0.1",
		},
	}, nil
}

// Initialized implements initialized method.
// https://microsoft.github.io/language-server-protocol/specification#initialized
func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	return nil
}

// Shutdown implements shutdown method.
// https://microsoft.github.io/language-server-protocol/specification#shutdown
func (s *Server) Shutdown(ctx context.Context) (err error) {
	ctx.Done()
	return nil
}

// Exit implements exit method.
// https://microsoft.github.io/language-server-protocol/specification#exit
func (s *Server) Exit(ctx context.Context) (err error) {
	ctx.Done()
	return nil
}
