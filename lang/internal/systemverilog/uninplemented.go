package systemverilog

import (
	"go.lsp.dev/jsonrpc2"
)

//
//func (s *Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
//	err = notImplemented("CodeAction")
//	return
//}
//
//func (s *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
//	err = notImplemented("CodeLens")
//	return
//}
//
//func (s *Server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
//	err = notImplemented("CodeLensResolve")
//	return
//}
//
//func (s *Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
//	err = notImplemented("ColorPresentation")
//	return
//}
//
//// Completion implements textDocument/completion method.
//// https://microsoft.github.io/language-server-protocol/specification#textDocument_completion
//func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
//	err = notImplemented("")
//	return
//	//return s.completion(ctx, params)
//}
//
//func (s *Server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
//	err = notImplemented("CompletionResolve")
//	return
//}
//
//func (s *Server) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
//	err = notImplemented("Declaration")
//	return
//}
//
//// Definition implements textDocument/definition method.
//// https://microsoft.github.io/language-server-protocol/specification#textDocument_definition
//func (s *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
//	err = notImplemented("")
//	return
//	//return s.definition(ctx, params)
//}
//
//// DidChange implements textDocument/didChange method.
//// https://microsoft.github.io/language-server-protocol/specification#textDocument_didChange
//func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
//	err = notImplemented("")
//	return
//	//return s.didChange(ctx, params)
//}
//
//func (s *Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
//	err = notImplemented("DidChangeConfiguration")
//	return
//}
//
//func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
//	err = notImplemented("DidChangeWatchedFiles")
//	return
//}
//
//// DidChangeWorkspaceFolders implements workspace/didChangeWorkspaceFolders method.
//// https://microsoft.github.io/language-server-protocol/specification#workspace_didChangeWorkspaceFolders
//func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
//	err = notImplemented("")
//	return
//	//return s.changeWorkspace(ctx, params.Event)
//}
//
//func (s *Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
//	err = notImplemented("DocumentColor")
//	return
//}
//
//func (s *Server) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
//	err = notImplemented("DocumentHighlight")
//	return
//}
//
//func (s *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
//	err = notImplemented("DocumentLink")
//	return
//}
//
//func (s *Server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
//	err = notImplemented("DocumentLinkResolve")
//	return
//}
//
//func (s *Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
//	err = notImplemented("DocumentSymbol")
//	return
//}
//
//func (s *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
//	err = notImplemented("ExecuteCommand")
//	return
//}
//
//func (s *Server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
//	err = notImplemented("FoldingRanges")
//	return
//}
//

//func (s *Server) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
//	err = notImplemented("Hover")
//	return
//}
//
//func (s *Server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
//	err = notImplemented("Implementation")
//	return
//}

//func (s *Server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
//	err = notImplemented("PrepareRename")
//	return
//}
//

//func (s *Handler) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
//	err = notImplemented("References")
//	return
//}
//
//func (s *Handler) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("Rename")
//	return
//}
//
//func (s *Handler) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
//	err = notImplemented("SignatureHelp")
//	return
//}
//
//func (s *Handler) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
//	err = notImplemented("Symbols")
//	return
//}
//
//func (s *Handler) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
//	err = notImplemented("TypeDefinition")
//	return
//}
//
//func (s *Handler) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
//	err = notImplemented("WillSave")
//	return
//}
//
//func (s *Handler) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
//	err = notImplemented("WillSaveWaitUntil")
//	return
//}
//
//func (s *Handler) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) CodeLensRefresh(ctx context.Context) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) SemanticTokensRefresh(ctx context.Context) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (s *Handler) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
//	err = notImplemented("")
//	return
//}
//
func notImplemented(method string) error {
	return jsonrpc2.Errorf(jsonrpc2.MethodNotFound, "method %q not implemented", method)
}
