@echo off
setlocal
cd /d "%~dp0"
for %%D in (01_basic_time 02_internal_packages 03_studentapp) do (
    echo Checking %%D
    go -C %%D test -count=1 ./...
    if errorlevel 1 goto failed
    go -C %%D vet ./...
    if errorlevel 1 goto failed
    go -C %%D build ./...
    if errorlevel 1 goto failed
)
echo All checks passed.
pause
exit /b 0
:failed
echo A check failed. Read the error above.
pause
exit /b 1
