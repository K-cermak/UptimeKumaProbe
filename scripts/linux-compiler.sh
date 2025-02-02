#!/bin/bash

API_DIR="$(pwd)/../api"
CLI_DIR="$(pwd)/../cli"
BIN_DIR="$(pwd)/../bin"

echo "API Directory: $API_DIR"
echo "CLI Directory: $CLI_DIR"
echo "BIN Directory: $BIN_DIR"

cd "$API_DIR" || {
    echo "Failed to change to API directory";
    exit 1;
}

GOOS=linux
GOARCH=amd64
go build -o kprobe-api || { 
    echo "Failed to build API application";
    exit 1;
}

[ ! -d "$BIN_DIR" ] && mkdir -p "$BIN_DIR" || {
    echo "Failed to create bin directory";
    exit 1;
}

mv -f kprobe-api "$BIN_DIR" || {
    echo "Failed to move API binary to bin directory";
    exit 1;
}

echo "[SUCCESS] API Build and move completed successfully"

cd "$CLI_DIR" || {
    echo "Failed to change to CLI directory";
    exit 1;
}

GOOS=linux
GOARCH=amd64
go build -o kprobe-cli || {
    echo "Failed to build CLI application";
    exit 1;
}

mv -f kprobe-cli "$BIN_DIR" || {
    echo "Failed to move CLI binary to bin directory";
    exit 1;
}

echo "[SUCCESS] CLI Build and move completed successfully"
