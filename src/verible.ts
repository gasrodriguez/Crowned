import * as vscode from 'vscode';
import * as child from 'child_process';
import * as path from 'path';

export function format(document: vscode.TextDocument, range: vscode.Range): vscode.TextEdit[] {
    const config = vscode.workspace.getConfiguration();
    const formatCommand = config.get("crowned.formatCommand");
    if (formatCommand === "") { return []; }
    let command = [];
    command.push(formatCommand);
    command.push('--lines ' + range.start.line + '-' + range.end.line);
    command.push(path.basename(document.uri.fsPath));
    const commandStr = command.join(' ').trim();
    console.debug(commandStr);
    try {
        const response = child.execSync(commandStr, {
            cwd: path.dirname(document.uri.fsPath),
        });
        return [new vscode.TextEdit(range, response.toString())];
    }
    catch (e: unknown) {
        console.error(e);
        return [];
    }
}

export async function lint(document: vscode.TextDocument): Promise<vscode.Diagnostic[]> {
    const config = vscode.workspace.getConfiguration();
    const lintCommand = config.get("crowned.lintCommand");
    if (lintCommand === "") { return []; }
    let command = [];
    command.push(lintCommand);
    command.push('--lint_fatal=false');
    command.push('--parse_fatal=false');
    command.push(path.basename(document.uri.fsPath));
    const commandStr = command.join(' ').trim();
    console.debug(commandStr);
    let diagnostics: vscode.Diagnostic[] = [];
    try {
        const response = child.execSync(commandStr, {
            cwd: path.dirname(document.uri.fsPath),
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

