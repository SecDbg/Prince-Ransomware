@echo off
setlocal enabledelayedexpansion

set "CURRENT_DIR=%~dp0"

cd /d "%CURRENT_DIR%Builder"

go build -ldflags "-s -w" -o "%CURRENT_DIR%Builder.exe" main.go
