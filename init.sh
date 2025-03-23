#!/bin/bash

TRACKGIT_PATH=$(pwd)

# Ensure ~/.bash_profile sources ~/.bashrc
if [ -f ~/.bash_profile ]; then
    if ! grep -q "source ~/.bashrc" ~/.bash_profile; then
        echo "source ~/.bashrc" >> ~/.bash_profile
    fi
else
    echo "source ~/.bashrc" > ~/.bash_profile
fi

# Export TRACKGIT_PATH in ~/.bashrc if not already present.
if ! grep -q "export TRACKGIT_PATH=" ~/.bashrc; then
    echo "export TRACKGIT_PATH='$TRACKGIT_PATH'" >> ~/.bashrc
else
    echo "TRACKGIT_PATH is already exported in bashrc"
fi

# Add alias 'trackgit' to ~/.bashrc if it doesn't exist.
chmod +x ./ContributionVisualizer/Scripts/trackgit.command.sh
if ! grep -q "alias trackgit=" ~/.bashrc; then
    echo "alias trackgit='./ContributionVisualizer/Scripts/trackgit.command.sh'" >> ~/.bashrc
    echo "Alias trackgit added to bashrc"
else
    echo "Alias trackgit already exists in bashrc"
fi

# Source the updated ~/.bashrc to load new environment variables and aliases.
source ~/.bashrc

# Set up Git hooks.
chmod +x ./CommitManager/Scripts/setup.post.hooks.sh
./CommitManager/Scripts/setup.post.hooks.sh

go build
go run main.go
