#!/bin/bash

export TRACKGIT_PATH=$(pwd)
chmod +x ./CommitManager/Scripts/setup.post.hooks.sh
./CommitManager/Scripts/setup.post.hooks.sh
go run main.go