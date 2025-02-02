:: DO NOT RUN THIS FILE BY RUN BUTTON IN VSCODE - IT MAY NOT WORK AS EXPECTED

@echo off

setlocal enabledelayedexpansion

set "API_DIR=%CD%\..\api"
set "CLI_DIR=%CD%\..\cli"
set "BIN_DIR=%CD%\..\bin"

echo API Directory: !API_DIR!
echo CLI Directory: !CLI_DIR!
echo BIN Directory: !BIN_DIR!

cd /d "!API_DIR!" || (
    echo Failed to change to API directory
    exit /b 1
)

set GOOS=linux
set GOARCH=amd64
go build -o api-server || (
    echo Failed to build Go application
    exit /b 1
)

if not exist "!BIN_DIR!" (
    mkdir "!BIN_DIR!" || (
        echo Failed to create bin directory
        exit /b 1
    )
)

move /Y api-server "!BIN_DIR!" || (
    echo Failed to move binary to bin directory
    exit /b 1
)

echo [SUCCESS] API Server Build and move completed successfully

cd /d "!CLI_DIR!" || (
    echo Failed to change to CLI directory
    exit /b 1
)

set GOOS=linux
set GOARCH=amd64
go build -o cli || (
    echo Failed to build Go application
    exit /b 1
)

move /Y cli "!BIN_DIR!" || (
    echo Failed to move binary to bin directory
    exit /b 1
)

echo [SUCCESS] CLI Build and move completed successfully

endlocal