@echo off
REM Migration script for Windows
REM Builds and runs the migration tool

echo Building migration tool...
go build -o ./build/migrate.exe ./cmd/migrate

if %errorlevel% neq 0 (
    echo Build failed!
    exit /b 1
)

echo Running migrations...
./build/migrate.exe

if %errorlevel% neq 0 (
    echo Migration failed!
    exit /b 1
)

echo Migration completed successfully!
