# GoModRename

GoModRename is a command-line tool designed to simplify the process of renaming Go modules. It automatically updates the module name in the go.mod file and adjusts all import paths in the project's Go files.

## Installation

To install GoModRename, use the following command:

```sh
go install github.com/hayletdomybest/gdr@v0.1.0-alpha
```

## Usage

```sh
gdr rename [new_name] -p  [project_path]
```

new_name: The new module name you want to use
project_path: (Optional) The path to your Go project. If not provided, the current directory will be used.
