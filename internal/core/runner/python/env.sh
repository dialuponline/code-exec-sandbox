#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <src> <dest>"
    exit 1
fi

src="$1"
dest="$2"

# Function to copy and link files
copy_and_link() {
    local src_file="$1"
    local dest_file="$2"

    if [ -L "$src_file" ]; then
        # If src_file is a symbolic link, copy it without changing permissions
        cp -P "$src_file" "$dest_file"
    elif [ -b "$src_file" ] || [ -c "$src_file" ]; then
        # If src_file is a device file, copy it and change permissions
        cp "$src_file" "$dest_file"
        ch