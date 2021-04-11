# Crowned

> This is the README for ~~your~~ my extension "crowned". After writing up a brief description, we recommend including the following sections.

## Features

> Describe specific features of your extension including screenshots of your extension in action. Image paths are relative to this README file.

OK, let's see: Crowned is a VS Code extension for Systemverilog and verilog formatting and linting using [verible](https://google.github.io/verible/). For everything else please use the amazing [SystemVerilog - Language Support](https://github.com/eirikpre/VSCode-SystemVerilog) extension.


> Tip: Many popular extensions utilize animations. **This is an excellent way to show off your extension!**

C'mon, that's the worst thing about most VS code extensions! Makes me go away as soon as possible without reading any possibly useful information...

## Requirements

Verible binaries. Download the right ones for your distro from [here](https://github.com/google/verible/releases), then extract and copy anywhere in the filesystem. I use Ubuntu-20.04 binaries running on WSL on Windows 10, and copied the linux binaries in a windows folder, but can also be installed in /usr/local/bin inside WSL.

## Extension Settings

This extension contributes the following settings:

* `crowned.lintCommand` = `verible-verilog-lint` (default)
* `crowned.formatCommand` = `verible-verilog-lint` (default)

On linux you may leave the defaults. On Windows/WSL a little trick is required. To get the WSL path of the verible bin directory:
```
$ wsl wslpath <path-to-verible-bin>
```
e.g.:
```
$ wsl wslpath 'C:\FpgaTools\verible-v0.0-1051-gd4cd328\bin'
/mnt/c/FpgaTools/verible-v0.0-1051-gd4cd328/bin
```
Then set the configuration settings as:
* `crowned.lintCommand` = `/mnt/c/FpgaTools/verible-v0.0-1051-gd4cd328/bin/verible-verilog-lint`
* `crowned.formatCommand` = `/mnt/c/FpgaTools/verible-v0.0-1051-gd4cd328/bin/verible-verilog-format`

## Known Issues

> Calling out known issues can help limit users opening duplicate issues against your extension.

Not known yet ;)

## Release Notes

> Users appreciate release notes as you update your extension.

Don't expect any.

**Enjoy!**
