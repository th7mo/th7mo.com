#!/bin/bash

# Set the paths
pages_directory="src/pages"
overview_file="src/pages/overview.astro"
exception_file="404-not-found"

# Get a list of all Astro files in src/pages/
astro_files=($(find "$pages_directory" -name "*.astro"))

# Extract the file names without the path and extension
astro_file_names=()
for file in "${astro_files[@]}"; do
  astro_file_names+=("$(basename "$file" .astro)")
done

# Check each Astro file against the overview.astro, excluding the exception file
for file_name in "${astro_file_names[@]}"; do
  if [[ "$file_name" != "$exception_file" ]]; then
    grep -q "$file_name" "$overview_file" || echo "Not listed in overview: $file_name"
  fi
done
