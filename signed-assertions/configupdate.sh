#!/bin/bash

# Set the root directory where you want to start searching for .cfg files
root_directory="/home/tom/Documents/signed-assertions/samples/anonymousMode"

# Define the old and new IP addresses
old_ip="192.168.1.107"
new_ip="10.10.20.32"

# Use the find command to search for .cfg files in subdirectories
find "$root_directory" -type f -name "*.cfg" -print0 | while IFS= read -r -d $'\0' config_file; do
    # Use sed to replace the old IP with the new IP in the file
    sed -i "s/$old_ip/$new_ip/g" "$config_file"
    echo "Updated IP in $config_file"
done
