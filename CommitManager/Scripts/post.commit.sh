#! /bin/bash

# Get the last commit hash
COMMIT_HASH=$(git rev-parse HEAD)

# Fetch the TRACKGIT_PATH from the environment
# This should be the path where the repository with your Go programs is located
if [ -z "$TRACKGIT_PATH" ]; then
    echo "TRACKGIT_PATH is not set. Please set TRACKGIT_PATH in your environment."
    exit 1
fi

go_program="$TRACKGIT_PATH/CommitManager/Hooks/PostCommit/post.commit.exec"

if [ ! -x "$go_program" ]; then
    echo "Git Tracker path incorrect, got $go_program"
    exit 1
fi

"$go_program" $COMMIT_HASH

