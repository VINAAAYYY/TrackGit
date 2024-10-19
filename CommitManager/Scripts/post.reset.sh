# #!/bin/bash
# source ~/.bashrc
# previous_head=$(git rev-parse HEAD)

# command git reset "$@"
# echo "reset hit"
# # Check if the reset was successful
# if [ $? -eq 0 ]; then
#     new_head=$(git rev-parse HEAD)
#     commits=$(git rev-list "$new_head".."$previous_head")

#     if [ -n "$commits" ]; then
#         if [ -z "$TRACKGIT_PATH" ]; then
#             echo "TRACKGIT_PATH is not set. Please set TRACKGIT_PATH in your environment."
#             exit 1
#         fi

#         go_program="$TRACKGIT_PATH/CommitManager/Hooks/PostCheckout/post.checkout.exec"
#         cd "$TRACKGIT_PATH/CommitManager/Hooks/PostCheckout" || exit
#         go build -o post.checkout.exec post.checkout.exec.go

#         if [ ! -x "$go_program" ]; then
#             echo "Git Tracker file path non-executable $go_program. Attempting to make it executable"
#             chmod +x "$go_program"
#         fi
#         "$go_program" "$commits"
#     fi
# else
#     echo "Failed to execute git reset hook."
#     exit 1
# fi
