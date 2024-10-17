#!/bin/bash

previous_head=$(git rev-parse ORIG_HEAD)
new_head=$(git rev-parse HEAD)

# Check if a reset, rebase, or similar operation changed the HEAD
if [ "$previous_head" != "$new_head" ]; then
    commits=$(git rev-list $new_head..$previous_head)

    if [ -n "$commits" ]; then
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

        "$go_program" "$commits"
    fi
fi
