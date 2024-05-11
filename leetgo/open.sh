#!/bin/bash

# echo $1 $2 $3
#{{.Folder}} {{.CodeFile}} {{.DescriptionFile}}

if [[ "$(echo $2 | grep 'solution.cpp')" != "" ]]; then
    cat .vscode/compile_commands.json | jq ".[].directory |= \"$1\"" > .vscode/_compile_commands.json
    mv .vscode/_compile_commands.json .vscode/compile_commands.json
    code "$2" "$3"
else
    code "$2" "$3"
fi
