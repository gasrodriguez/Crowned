import * as vscode from 'vscode';
import * as vscodeClient from 'vscode-languageclient';

const extensionDisplayName = 'Crowned';

let client: vscodeClient.LanguageClient;

export function activate(context: vscode.ExtensionContext) {
    const serverCmd: string = vscode.workspace.getConfiguration().get("crowned.serverCommand")!;
    const serverEnv: any = vscode.workspace.getConfiguration().get("crowned.serverEnv");

    const run: vscodeClient.Executable = {
        command: serverCmd,
        options: {
            env: serverEnv
        }
    };

    // If the extension is launched in debug mode then the debug server options are used
    // Otherwise the run options are used
    let serverOptions: vscodeClient.ServerOptions = {
        run,
        debug: run,
    };

    // Options to control the language client
    let clientOptions: vscodeClient.LanguageClientOptions = {
        // Register the server for plain text documents
        documentSelector: [
            {
                scheme: "file", language: "systemverilog"
            },
            {
                scheme: "file", language: "verilog"
            }
        ],
    };

    // Create the language client and start the client.
    client = new vscodeClient.LanguageClient(
        extensionDisplayName,
        extensionDisplayName,
        serverOptions,
        clientOptions
    );

    // Start the client. This will also launch the server
    client.start();

    // Register command to restart server
    context.subscriptions.push(vscode.commands.registerCommand('crowned.serverRestart', serverRestart));
}

function serverRestart() {
    client.stop().then(() => client.start());
}

// This method is called when your extension is deactivated
export function deactivate() { client.stop(); }
