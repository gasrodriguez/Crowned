import * as vscode from 'vscode';
import * as child from 'child_process';
import * as path from 'path';
import * as fs from 'fs';

const workingDir = ((document: vscode.TextDocument) => {
    let cwd = path.dirname(document.uri.fsPath);
    const wsFolder = vscode.workspace.getWorkspaceFolder(document.uri);
    if (wsFolder) {
        cwd = wsFolder.uri.fsPath;
    }
    return cwd;
});

const filePath = ((cwd: string, document: vscode.TextDocument) => {
    const relPath = path.relative(cwd, document.uri.fsPath);
    return relPath.replace(/\\/g, '/');
});

const configFile = ((cwd: string, filename: string) => {
    const fsPath = path.join(cwd, '.verible', filename);
    if (fs.existsSync(fsPath)) {
        return ".verible/" + filename;
    } else {
        return "";
    }
});

const logCommand = ((cwd: string, command: string) => {
    console.debug(cwd + '> ' + command);
});

export function format(document: vscode.TextDocument, _options: vscode.FormattingOptions, _token: vscode.CancellationToken): vscode.TextEdit[] {
    let cwd = workingDir(document);
    const config = vscode.workspace.getConfiguration();
    const formatCommand = config.get("crowned.formatCommand");
    if (formatCommand === "") { return []; }
    let command = [];
    command.push(formatCommand);
    command.push('--flagfile=' + configFile(cwd, 'format.flags'));
    command.push(filePath(cwd, document));
    try {
        const commandStr = command.join(' ').trim();
        logCommand(cwd, commandStr);
        const response = child.execSync(commandStr, {
            cwd: cwd,
        });
        const range = new vscode.Range(document.lineAt(0).range.start,
            document.lineAt(document.lineCount - 1).range.end);
        return [new vscode.TextEdit(range, response.toString())];
    }
    catch (e: unknown) {
        console.error(e);
        return [];
    }
}

export async function lint(document: vscode.TextDocument, outputChannel: vscode.OutputChannel): Promise<vscode.Diagnostic[]> {
    let cwd = workingDir(document);
    const config = vscode.workspace.getConfiguration();
    const lintCommand = config.get("crowned.lintCommand");
    if (lintCommand === "") { return []; }
    let command = [];
    command.push(lintCommand);
    command.push('--lint_fatal=false');
    command.push('--parse_fatal=false');
    // command.push('--flagfile=' + configFile(cwd, 'lint.flags'));
    // command.push('--rules_config=' + configFile(cwd, 'lint.rules'));
    // command.push('--waiver_files=' + configFile(cwd, 'lint.waiver'));
    command.push(filePath(cwd, document));
    let diagnostics: vscode.Diagnostic[] = [];
    try {
        const commandStr = command.join(' ').trim();
        logCommand(cwd, commandStr);
        outputChannel.appendLine(cwd + '> ' + commandStr);
        const response = child.execSync(commandStr, {
            cwd: cwd,
        });
        const lines = response.toString().split('\n');
        lines.forEach((line, i) => {
            const [_, lineStr, colStr, ...message] = line.split(':');
            const lineNum = Number(lineStr) - 1;
            const colNum = Number(colStr) - 1;
            const messageStr = message.join(':');
            let severity = vscode.DiagnosticSeverity.Error;
            if (message && messageStr.search('error') === -1) {
                severity = vscode.DiagnosticSeverity.Warning;
            }
            diagnostics.push({
                severity: severity,
                range: new vscode.Range(lineNum, colNum, lineNum, Number.MAX_VALUE),
                message: messageStr,
                code: i,
                source: 'verible-verilog-lint'
            });
        });
    }
    catch (e: unknown) {
        console.error(e);
    }
    return diagnostics;
}

