#!/bin/bash

 comm -23 <(ls src/*.html | sed 's/src\///' | sort) <(grep -o '<a href="[^"]*' src/overview.html | sed 's/<a href="//' | sort) | sed 's/^/Missing file: /'
