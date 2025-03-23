#!/bin/bash

# Function to call the calender.flex.service
call_service() {
    echo "*** Visualizing your contributions for the last $1... ***"
    source ~/.bashrc
    if [ -z "$TRACKGIT_PATH" ]; then
        echo "TRACKGIT_PATH is not set. Please set TRACKGIT_PATH in your environment."
        exit 1
    fi

    go_program="$TRACKGIT_PATH/ContributionVisualizer/InitService/main.exec"

    cd "$TRACKGIT_PATH/ContributionVisualizer/InitService" || exit
    go build -o main.exec main.go
    
    if [ ! -x "$go_program" ]; then
        echo "Git Tracker file path non-executable $go_program. Attempting to make it executable"
        chmod +x "$go_program"
    fi
    "$go_program" "$1"
}

# Check the arguments
if [ "$1" == "visualize" ]; then
    case "$2" in
        --week)
            call_service "week"
            ;;
        --month)
            call_service "month"
            ;;
        --year)
            call_service "year"
            ;;
        *)
            echo "Invalid option. Use --week, --month, or --year."
            ;;
    esac
else
    echo "Invalid Command. Usage: trackgit visualize --week|--month|--year"
fi