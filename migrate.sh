#!/bin/bash
# Migration script for Linux/Mac
# Builds and runs the migration tool

echo "Building migration tool..."
go build -o ./build/migrate ./cmd/migrate

if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Running migrations..."
./build/migrate

if [ $? -ne 0 ]; then
    echo "Migration failed!"
    exit 1
fi

echo "Migration completed successfully!"
