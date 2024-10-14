#! /bin/bash

# Get the last commit hash
COMMIT_HASH=$(git rev-parse HEAD)

# path to go program relative to .git/hooks. HOW?
.go/file/path $COMMIT_HASH
