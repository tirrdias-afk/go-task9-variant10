@echo off
chcp 65001 >nul
cd /d "%~dp0"
:menu
echo.
echo 1 - Basic task 10: current time
echo 2 - Intermediate task 10: internal packages
echo 3 - Laboratory task 10: StudentApp
echo 4 - StudentApp automatic demo
echo 0 - Exit
set "task_choice="
set /p "task_choice=Choose: "
if "%task_choice%"=="0" exit /b
if "%task_choice%"=="1" goto basic
if "%task_choice%"=="2" goto internal
if "%task_choice%"=="3" goto app
if "%task_choice%"=="4" goto demo
goto menu
:basic
pushd 01_basic_time
basic10.exe
popd
pause
goto menu
:internal
pushd 02_internal_packages
internal10.exe
popd
pause
goto menu
:app
pushd 03_studentapp
studentapp.exe
popd
pause
goto menu
:demo
pushd 03_studentapp
studentapp.exe -demo
popd
pause
goto menu
