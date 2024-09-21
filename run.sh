#!/bin/bash

# Set environment variables
export DB_HOST=localhost
export DB_USERNAME=your_username
export DB_PASSWORD=your_password
export DB_NAME=your_database
export DB_PORT=5432
export DB_SSL_MODE=disable
export DB_SCHEMA=public

# Enable CGO and build the application
export CGO_ENABLED=1
go build -o gorm_migrator ./cmd/main.go

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed"
    exit 1
fi

# Run the application
echo "Running gorm_migrator..."
./gorm_migrator
