@echo off
REM BettyPy Installer for Windows

set INSTALL_DIR=%ProgramFiles%\BettyPy
set SCRIPT_NAME=bettypy

net session >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Please run this script as Administrator.
    pause
    exit /b 1
)

echo [INFO] Creating installation directory...
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

echo [INFO] Copying files...
copy /y "%~dp0%SCRIPT_NAME%" "%INSTALL_DIR%\%SCRIPT_NAME%"

echo [INFO] Adding BettyPy to system PATH...
setx PATH "%PATH%;%INSTALL_DIR%" /M

echo [INFO] Verifying installation...
if exist "%INSTALL_DIR%\%SCRIPT_NAME%" (
    echo [SUCCESS] BettyPy has been installed successfully!
    echo You can now use the "bettypy" command in your terminal.
) else (
    echo [ERROR] Installation failed. Please check the script and try again.
)

pause