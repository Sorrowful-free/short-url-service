@echo off
echo Collecting memory profile...
go tool pprof -proto -output=profiles\result.pprof http://localhost:6060/debug/pprof/heap
if %ERRORLEVEL% EQU 0 (
    echo Memory profile saved to profiles\result.pprof
) else (
    echo Failed to collect memory profile. Make sure the application is running and pprof server is accessible at http://localhost:6060
    exit /b 1
)

