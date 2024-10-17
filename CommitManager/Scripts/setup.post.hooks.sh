#!/bin/bash

GLOBAL_HOOKS_DIR="$HOME/.global-git-hooks"

# Create the global hooks directory if not exists
if [ ! -d "$GLOBAL_HOOKS_DIR" ]; then
    mkdir -p "$GLOBAL_HOOKS_DIR"
fi

cp ./post.commit.sh "$GLOBAL_HOOKS_DIR/post-commit"
cp ./post.rewrite.sh "$GLOBAL_HOOKS_DIR/post-rewrite"
cp ./post.checkout.sh "$GLOBAL_HOOKS_DIR/post-checkout"

# Ensuring the hooks are executable
chmod +x "$GLOBAL_HOOKS_DIR/post-commit"
chmod +x "$GLOBAL_HOOKS_DIR/post-rewrite"
chmod +x "$GLOBAL_HOOKS_DIR/post-checkout"

git config --global core.hooksPath "$GLOBAL_HOOKS_DIR"

echo "***** Global Git hooks configured successfully *****"
