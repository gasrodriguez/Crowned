import * as vscode from 'vscode';
import * as verible from './verible';
// import * as VHDLFormatter from './VHDLFormatter';

const extensionDisplayName = 'crowned';

const languageIdVhdl = 'vhdl';
const languageIdVerilog = 'verilog';
const languageIdSystemVerilog = 'systemverilog';

let diagnosticCollection: vscode.DiagnosticCollection;

export function activate(context: vscode.ExtensionContext) {
    console.log('Crowned extension activated.');

    // context.subscriptions.push(vscode.languages.registerDocumentRangeFormattingEditProvider(
    //     languageIdVhdl, { provideDocumentRangeFormattingEdits: VHDLFormatter.format }
    // ));

    context.subscriptions.push(vscode.languages.registerDocumentFormattingEditProvider(
        languageIdVerilog, { provideDocumentFormattingEdits: verible.format }
    ));

    context.subscriptions.push(vscode.languages.registerDocumentFormattingEditProvider(
        languageIdSystemVerilog, { provideDocumentFormattingEdits: verible.format }
    ));

    diagnosticCollection = vscode.languages.createDiagnosticCollection(extensionDisplayName);
    context.subscriptions.push(vscode.commands.registerCommand('crowned.clear_diagnostics', commandClearDiagnostics));

    context.subscriptions.push(
        vscode.workspace.onDidOpenTextDocument(didSaveTextDocument),
        vscode.workspace.onDidSaveTextDocument(didSaveTextDocument),
        vscode.workspace.onDidCloseTextDocument(didCloseTextDocument),
        vscode.workspace.onDidRenameFiles(didRenameFiles),
    );
}

function didSaveTextDocument(document: vscode.TextDocument) {
    let promise: Promise<vscode.Diagnostic[]>;
    switch (document.languageId) {
        case languageIdVerilog:
        case languageIdSystemVerilog:
            promise = verible.lint(document);
            promise
                .then((diagnostics: vscode.Diagnostic[]) => {
                    diagnosticCollection.delete(document.uri);
                    diagnosticCollection.set(document.uri, diagnostics);
                })
                .catch((e) => {
                    console.error(e);
                });
            break;
    }
}

function didCloseTextDocument(document: vscode.TextDocument) {
    diagnosticCollection.delete(document.uri);
}

function didRenameFiles(e: vscode.FileRenameEvent) {
    e.files.forEach(element => {
        diagnosticCollection.delete(element.oldUri);
    });
}

function commandClearDiagnostics() {
    diagnosticCollection.clear();
}

// this method is called when your extension is deactivated
export function deactivate() { }
