#!/bin/bash
source ~/.bashrc
GLOBAL_HOOKS_DIR="$HOME/.global-git-hooks"

echo "***** Setting up global git hooks *****"

# Create the global hooks directory if not exists
if [ ! -d "$GLOBAL_HOOKS_DIR" ]; then
    mkdir -p "$GLOBAL_HOOKS_DIR"
fi

cp "$TRACKGIT_PATH/CommitManager/Scripts/post.commit.sh" "$GLOBAL_HOOKS_DIR/post-commit"
# cp "$TRACKGIT_PATH/CommitManager/Scripts/post.checkout.sh" "$GLOBAL_HOOKS_DIR/post-checkout"
# cp "$TRACKGIT_PATH/CommitManager/Scripts/post.reset.sh" "$GLOBAL_HOOKS_DIR/post-reset"

chmod +x "$GLOBAL_HOOKS_DIR/post-commit"
chmod +x "$GLOBAL_HOOKS_DIR/post-rewrite"
chmod +x "$GLOBAL_HOOKS_DIR/post-checkout"
# chmod +x "$GLOBAL_HOOKS_DIR/post-reset"

# git config --global alias.reset "$GLOBAL_HOOKS_DIR/post-reset"
git config --global core.hooksPath "$GLOBAL_HOOKS_DIR"

echo "***** Global Git hooks configured successfully *****"