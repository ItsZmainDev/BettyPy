#!/bin/bash

# Variables
INSTALL_DIR="/usr/local/bin"
MAN_DIR="/usr/local/share/man/man1"
SCRIPT_NAME="bettypy"
MAN_PAGE_NAME="bettypy.1"
GO_VERSION="1.20.5"
GO_TAR="go$GO_VERSION.linux-amd64.tar.gz"
GO_URL="https://golang.org/dl/$GO_TAR"

# Check for root permissions
if [ "$EUID" -ne 0 ]; then
  echo "Please run this script as root or with sudo."
  exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
  echo "Go is not installed. Installing Go..."
  
  # Download and install Go
  wget "$GO_URL" -O "/tmp/$GO_TAR"
  tar -C /usr/local -xzf "/tmp/$GO_TAR"
  rm "/tmp/$GO_TAR"
  
  # Add Go to PATH
  export PATH=$PATH:/usr/local/go/bin
  echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
  source ~/.bashrc
  
  echo "Go has been installed successfully."
else
  echo "Go is already installed."
fi

# Clean previous installation
echo "Cleaning previous installations..."
if [ -f "$INSTALL_DIR/$SCRIPT_NAME" ]; then
  rm -f "$INSTALL_DIR/$SCRIPT_NAME"
  echo "Old binary removed from $INSTALL_DIR."
fi
if [ -f "$MAN_DIR/$MAN_PAGE_NAME" ]; then
  rm -f "$MAN_DIR/$MAN_PAGE_NAME"
  echo "Old man page removed from $MAN_DIR."
fi

# Build the Go binary
echo "Building the Go binary..."
go build -o "$SCRIPT_NAME" ./cmd/bettypy

# Install the binary
echo "Installing the $SCRIPT_NAME command..."
mv "$SCRIPT_NAME" "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/$SCRIPT_NAME"

# Create the man page
echo "Installing the man page..."
mkdir -p "$MAN_DIR"
cat << EOF > "$MAN_DIR/$MAN_PAGE_NAME"
.TH bettypy 1 "12 April 2025" "BettyPy v1.0" "Man page for bettypy"
.SH NAME
bettypy \- Python style checker for Betty
.SH SYNOPSIS
.B bettypy
.I python_file
.SH DESCRIPTION
BettyPy is a Python style checker tool based on Betty's style rules.
.SH OPTIONS
No specific options are required. Simply provide the path to the Python file to analyze.
.SH EXAMPLES
Usage example:
.PP
.B bettypy
.I example.py
.SH AUTHORS
Martin Clement \- https://github.com/ItsZmainDev
EOF

# Update the man database
echo "Updating the man database..."
mandb

# Finish installation
echo "Installation complete. You can now use the '$SCRIPT_NAME' command and view the man page with 'man $SCRIPT_NAME'."