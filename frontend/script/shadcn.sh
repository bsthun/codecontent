#!/bin/bash

DIR=$(dirname $0)
SHADCN_FILE="$DIR/shadcn.txt"

if [ ! -f "$SHADCN_FILE" ]; then
    echo "Error: $SHADCN_FILE not found"
    exit 1
fi

# sort component names alphabetically
sort -o "$SHADCN_FILE" "$SHADCN_FILE"

# install each component
while IFS= read -r component; do
    # Skip empty lines and comments
    if [ -n "$component" ] && [[ ! "$component" =~ ^# ]]; then
        echo "Installing $component..."
        bun x shadcn-svelte@latest add "$component" < /dev/tty
    fi
done < "$SHADCN_FILE"

exit 0