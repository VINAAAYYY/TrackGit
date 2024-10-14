#!/bin/bash

# cases: git amend(replaced by a new commit), git rebase(rebased onto a new base commit)

# Accumulate the old commit hashes
affected_hashes=()

while read old_hash new_hash; do
    affected_hashes+=("$old_hash")
done

# path to go program relative to .git/hooks. HOW?
if [ ${#affected_hashes[@]} -gt 0 ]; then
    .go/file/path "${affected_hashes[@]}"
fi
