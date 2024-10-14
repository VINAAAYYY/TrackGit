#!/bin/bash

#cases: git reset --hard etc

# Get the previous HEAD state
previous_head=$(git rev-parse ORIG_HEAD)
new_head=$(git rev-parse HEAD)

# check if it is NOT cherry-pick or similar command
if [ "$previous_head" != "$new_head" ]; then
    # Get all commits between old and new heads
    commits=$(git rev-list $new_head..$previous_head)

    # path to go program relative to .git/hooks. HOW?
    if [ -n "$commits" ]; then
        .go/file/path $commits
    fi
fi
