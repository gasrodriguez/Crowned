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

    context.subscriptions.push(vscode.languages.registerDocumentRangeFormattingEditProvider(
        languageIdVerilog, { provideDocumentRangeFormattingEdits: verible.format }
    ));

    context.subscriptions.push(vscode.languages.registerDocumentRangeFormattingEditProvider(
        languageIdSystemVerilog, { provideDocumentRangeFormattingEdits: verible.format }
    ));

    diagnosticCollection = vscode.languages.createDiagnosticCollection(extensionDisplayName);

    context.subscriptions.push(
        vscode.workspace.onDidSaveTextDocument(didSaveTextDocument),
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

// this method is called when your extension is deactivated
export function deactivate() { }
