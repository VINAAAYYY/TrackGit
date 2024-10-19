#!/bin/bash
source ~/.bashrc

# cases: git amend(replaced by a new commit), git rebase(rebased onto a new base commit)
affected_hashes=()

while read old_hash new_hash; do
    affected_hashes+=("$old_hash")
done

if [ ${#affected_hashes[@]} -gt 0 ]; then
    if [ -z "$TRACKGIT_PATH" ]; then
        echo "TRACKGIT_PATH is not set. Please set TRACKGIT_PATH in your environment."
        exit 1
    fi

    go_program="$TRACKGIT_PATH/CommitManager/Hooks/PostCheckout/post.checkout.exec"
    cd "$TRACKGIT_PATH/CommitManager/Hooks/PostCheckout" || exit
    go build -o post.checkout.exec post.checkout.exec.go

    if [ ! -x "$go_program" ]; then
        echo "Git Tracker file path non-executable $go_program. Attempting to make it executable"
        chmod +x "$go_program"
    fi

    "$go_program" "$affected_hashes"
fi
