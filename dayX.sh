#!/bin/bash

# Check if a number is provided as an argument
if [ -z "$1" ]; then
  echo "Please provide a number as an argument."
  exit 1
fi

# Create the directory day<number>
DIR_NAME="day$1"
mkdir "$DIR_NAME"

# Change into the directory
cd "$DIR_NAME" || exit

# Create the main.go file
touch main.go

# Write content to the main.go file
cat <<EOL >> main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOL

# Open the current directory in VS Code
code .

# Print success message
echo "Directory $DIR_NAME created and main.go file added."
