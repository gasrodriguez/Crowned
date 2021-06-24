package systemverilog

import (
	"context"
	"fmt"
	"github.com/gasrodriguez/crowned/pkg/lsp"
	"go.lsp.dev/protocol"
)

const (
	ServerName = "Crowned SystemVerilog Language Server"
)

var (
	ServerVersion = "0.0.0"
)

type Handler struct {
	lsp.Handler
	workspacePath string
	files         *files
}

func NewHandler() *Handler {
	h := &Handler{}
	h.files = NewFiles()
	return h
}

func (o *Handler) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	if o.files.isModified(params.TextDocument.URI.Filename()) {
		// User must turn auto-save on so files won't be in state modified for long
		return nil, protocol.ErrContentModified
	}
	return o.doFormatting(ctx, params)
}

// DidOpen implements textDocument/didOpen method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didOpen
func (o *Handler) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	go o.publishDiagnostics(ctx, params.TextDocument.URI)
	return
}

// DidSave implements textDocument/didSave method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didSave
func (o *Handler) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	o.files.SavedFile(params.TextDocument.URI.Filename())
	go o.publishDiagnostics(ctx, params.TextDocument.URI)
	return
}

// DidClose implements textDocument/didClose method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didClose
func (o *Handler) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	return o.clearDiagnostics(ctx, params)
}

// Initialize implements initialize method.
// https://microsoft.github.io/language-server-protocol/specification#initialize
func (o *Handler) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	return o.doInitialize(ctx, params)
}

// Initialized implements initialized method.
// https://microsoft.github.io/language-server-protocol/specification#initialized
func (o *Handler) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	return o.doInitialized(ctx, params)
}

// Shutdown implements shutdown method.
// https://microsoft.github.io/language-server-protocol/specification#shutdown
func (o *Handler) Shutdown(_ context.Context) (err error) {
	o.ShowMessage(fmt.Sprintf("%s shutdown.", ServerName))
	return
}

func (o *Handler) DidCreateFiles(_ context.Context, params *protocol.CreateFilesParams) (err error) {
	for _, file := range params.Files {
		o.files.AddFile(protocol.URI(file.URI).Filename())
	}
	return
}

func (o *Handler) DidRenameFiles(_ context.Context, params *protocol.RenameFilesParams) (err error) {
	for _, file := range params.Files {
		o.files.RemoveFile(protocol.URI(file.OldURI).Filename())
		o.files.AddFile(protocol.URI(file.NewURI).Filename())
	}
	return
}

// DidChange implements textDocument/didChange method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didChange
func (o *Handler) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	o.files.ChangeFile(params.TextDocument.URI.Filename())
	return
}

//
//func (o *Handler) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
//	err = notImplemented("CodeAction")
//	return
//}
//
//func (o *Handler) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
//	err = notImplemented("CodeLens")
//	return
//}
//
//func (o *Handler) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
//	err = notImplemented("CodeLensResolve")
//	return
//}
//
//func (o *Handler) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
//	err = notImplemented("ColorPresentation")
//	return
//}
//
//// Completion implements textDocument/completion method.
//// https://microsoft.github.io/language-server-protocol/specification#textDocument_completion
//func (o *Handler) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
//	err = notImplemented("")
//	return
//	//return s.completion(ctx, params)
//}
//
//func (o *Handler) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
//	err = notImplemented("CompletionResolve")
//	return
//}
//
//func (o *Handler) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
//	err = notImplemented("Declaration")
//	return
//}
//
//// Definition implements textDocument/definition method.
//// https://microsoft.github.io/language-server-protocol/specification#textDocument_definition
//func (o *Handler) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
//	err = notImplemented("")
//	return
//	//return s.definition(ctx, params)
//}

func (o *Handler) DidDeleteFiles(_ context.Context, params *protocol.DeleteFilesParams) (err error) {
	for _, file := range params.Files {
		o.files.RemoveFile(protocol.URI(file.URI).Filename())
	}
	return
}

//func (o *Handler) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
//	err = notImplemented("DidChangeConfiguration")
//	return
//}
//
//func (o *Handler) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
//	err = notImplemented("DidChangeWatchedFiles")
//	return
//}
//
//// DidChangeWorkspaceFolders implements workspace/didChangeWorkspaceFolders method.
//// https://microsoft.github.io/language-server-protocol/specification#workspace_didChangeWorkspaceFolders
//func (o *Handler) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
//	err = notImplemented("")
//	return
//	//return s.changeWorkspace(ctx, params.Event)
//}
//
//func (o *Handler) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
//	err = notImplemented("DocumentColor")
//	return
//}
//
//func (o *Handler) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
//	err = notImplemented("DocumentHighlight")
//	return
//}
//
//func (o *Handler) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
//	err = notImplemented("DocumentLink")
//	return
//}
//
//func (o *Handler) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
//	err = notImplemented("DocumentLinkResolve")
//	return
//}
//
//func (o *Handler) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
//	err = notImplemented("DocumentSymbol")
//	return
//}
//
//func (o *Handler) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
//	err = notImplemented("ExecuteCommand")
//	return
//}
//
//func (o *Handler) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
//	err = notImplemented("FoldingRanges")
//	return
//}
//

//func (o *Handler) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
//	err = notImplemented("Hover")
//	return
//}
//
//func (o *Handler) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
//	err = notImplemented("Implementation")
//	return
//}

//func (o *Handler) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
//	err = notImplemented("PrepareRename")
//	return
//}
//

//func (o *Handler) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
//	err = notImplemented("References")
//	return
//}
//
//func (o *Handler) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("Rename")
//	return
//}
//
//func (o *Handler) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
//	err = notImplemented("SignatureHelp")
//	return
//}
//
//func (o *Handler) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
//	err = notImplemented("Symbols")
//	return
//}
//
//func (o *Handler) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
//	err = notImplemented("TypeDefinition")
//	return
//}
//
//func (o *Handler) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
//	err = notImplemented("WillSave")
//	return
//}
//
//func (o *Handler) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
//	err = notImplemented("WillSaveWaitUntil")
//	return
//}
//
//func (o *Handler) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("")
//	return
//}

// Exit implements exit method.
// https://microsoft.github.io/language-server-protocol/specification#exit
func (o *Handler) Exit(ctx context.Context) (err error) {
	o.ShowMessage(fmt.Sprintf("%s exited.", ServerName))
	return nil
}

//func (o *Handler) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("")
//	return
//}

//func (o *Handler) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
//	err = notImplemented("")
//	return
//}

//func (o *Handler) CodeLensRefresh(ctx context.Context) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) SemanticTokensRefresh(ctx context.Context) (err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
//	err = notImplemented("")
//	return
//}
//
//func (o *Handler) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
//	err = notImplemented("OnTypeFormatting")
//	return
//}
//
//func (o *Handler) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
//	err = notImplemented("RangeFormatting")
//	return
//}
//
