#!/bin/bash

# echo $1 $2 $3
#{{.Folder}} {{.CodeFile}} {{.DescriptionFile}}

if [[ "$(echo $2 | grep 'solution.cpp')" != "" ]]; then
    old_dir=$(jq -r '.[].directory' .vscode/compile_commands.json)
    sed -i s%$old_dir%$1%g .vscode/compile_commands.json
    code $2 $3
else
    code $2 $3
fi
