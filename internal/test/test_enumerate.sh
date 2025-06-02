#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

go run $SCRIPT_DIR/../../main.go

if [ $? -eq 0 ]; then
    echo "SUCCESS: enumerate ran without error."
else
    echo "FAIL: enumerate returned non-zero return code."
    exit 1
fi
