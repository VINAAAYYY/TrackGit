#! /bin/bash
source ~/.bashrc
COMMIT_HASH=$(git rev-parse HEAD)
REPO_PATH=$(git rev-parse --show-toplevel)
if [ -z "$TRACKGIT_PATH" ]; then
    echo "TRACKGIT_PATH is not set. Please set TRACKGIT_PATH in your environment."
    exit 1
fi

go_program="$TRACKGIT_PATH/CommitManager/Hooks/PostCommit/post.commit.exec"

cd "$TRACKGIT_PATH/CommitManager/Hooks/PostCommit" || exit
go build -o post.commit.exec post.commit.exec.go

if [ ! -x "$go_program" ]; then
    echo "Git Tracker file path non-executable $go_program. Attempting to make it executable"
    chmod +x "$go_program"
fi

"$go_program" "$COMMIT_HASH" "$REPO_PATH"

