@echo off
setlocal
cd /d "%~dp0"
where go >nul 2>nul
if errorlevel 1 (
    echo Go not found. Install Go 1.22 or newer from https://go.dev/dl/ and reopen this file.
    pause
    exit /b 1
)
go -C 01_basic_time build -trimpath -o basic10.exe .
if errorlevel 1 goto failed
go -C 02_internal_packages build -trimpath -o internal10.exe .
if errorlevel 1 goto failed
go -C 03_studentapp build -mod=vendor -trimpath -o studentapp.exe .
if errorlevel 1 goto failed
echo All three programs were built. Open START.cmd.
pause
exit /b 0
:failed
echo Build failed. Read the error above.
pause
exit /b 1
