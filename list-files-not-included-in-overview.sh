#!/bin/bash

ls src/*.html | grep -Ev '404-not-found.html|template.html|overview.html' | sed 's/src\///' | sort > /tmp/all_files
grep -o '<a href="[^"]*' src/overview.html | grep -Ev '404-not-found.html|template.html|overview.html' | sed 's/<a href="//' | sort > /tmp/overview_files

missing_files=$(comm -23 /tmp/all_files /tmp/overview_files)

if [ -n "$missing_files" ]; then
  echo "Missing files:"
  echo "$missing_files" | sed 's/^/  - /'
else
  echo "All HTML files are listed in src/overview.html"
fi

# Clean up temporary files
rm /tmp/all_files /tmp/overview_files
