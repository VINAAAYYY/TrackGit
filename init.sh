#!/bin/bash

TRACKGIT_PATH=$(pwd)

if [ -f ~/.bash_profile ]; then
    if ! grep -q "source ~/.bashrc" ~/.bash_profile; then
        echo "source ~/.bashrc" >> ~/.bash_profile
    fi
else
    echo "source ~/.bashrc" > ~/.bash_profile
fi

if ! grep -q "export TRACKGIT_PATH=" ~/.bashrc; then
    echo "export TRACKGIT_PATH='$TRACKGIT_PATH'" >> ~/.bashrc
else
    echo "TRACKGIT_PATH is already exported in ~/.bashrc"
fi

source ~/.bashrc

chmod +x ./CommitManager/Scripts/setup.post.hooks.sh
./CommitManager/Scripts/setup.post.hooks.sh
go build
go run main.go