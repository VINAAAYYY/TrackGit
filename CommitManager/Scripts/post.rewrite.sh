#!/bin/bash

# cases: git amend(replaced by a new commit), git rebase(rebased onto a new base commit)

# Accumulate the old commit hashes
affected_hashes=()

while read old_hash new_hash; do
    affected_hashes+=("$old_hash")
done

# path to go program relative to .git/hooks
if [ ${#affected_hashes[@]} -gt 0 ]; then
    # Fetch the TRACKGIT_PATH from the environment
    # This should be the path where the repository with your Go programs is located
    if [ -z "$TRACKGIT_PATH" ]; then
        echo "TRACKGIT_PATH is not set. Please set TRACKGIT_PATH in your environment."
        exit 1
    fi

    go_program="$TRACKGIT_PATH/CommitManager/Hooks/PostCheckout/post.checkout.exec"

    if [ ! -x "$go_program" ]; then
        echo "Git Tracker path incorrect, got $go_program"
        exit 1
    fi

    "$go_program" "$affected_hashes"
fi
