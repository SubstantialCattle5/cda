#!/bin/bash

# usage message
usage() {
    echo "Usage: $0 <operation> <pool-name> <image-name>"
    echo "  operation: 'export' for full export, 'export-diff' for export-diff"
    exit 1
}

# Function to check if rbd command exists
check_rbd_command() {
    if ! command -v rbd &>/dev/null; then
        echo "Error: 'rbd' command not found. Please make sure Ceph utilities are installed and accessible."
        exit 1
    fi
}

# Calculate the splits
calculate_splits() {
    local poolname="$1"
    local imagename="$2"

    # Fetch the size from rbd info
    size=10
    if [[ -z "$size" ]]; then
        echo "Error: Could not determine size of the image. Exiting."
        exit 1
    fi

    # Calculate 10% of the size
    local percent=$(echo "scale=0; $size * 0.1 / 1" | bc)

    echo "$percent"
}

#! TODO: checkout for overflow error
# Array to hold individual chunk hashes
chunk_hashes=()

# Function to calculate the combined checksum
calculate_combined_checksum() {
    combined_hash_input=$(printf "%s" "${chunk_hashes[@]}")
    echo -n "$combined_hash_input" | sha256sum
}

export_and_checksums() {
    local operation="$1"
    local poolname="$2"
    local imagename="$3"
    local splitsize=$(calculate_splits "$poolname" "$imagename")

    # Export RBD image, split into chunks, and hash each chunk
    snapshot_file="./GD.iso"
    split -b ${splitsize}G "${snapshot_file}" --filter='sha256sum' | (
        while read -r hash filename; do
            chunk_hashes+=("$hash")
        done
    )
}

# Main script logic

if [ "$#" -ne 3 ]; then
    usage
fi

operation="$1"
poolname="$2"
imagename="$3"

# Check for rbd installation

# # Run the export and calculate checksums

export_and_checksums "$operation" "$poolname" "$imagename"
calculate_combined_checksum
