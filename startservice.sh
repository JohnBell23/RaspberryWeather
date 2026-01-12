#!/bin/bash

# Stop the service if it's running
sudo systemctl stop RaspberryWeather 2>/dev/null || true

# Create the installation directory
mkdir -p ~/.RaspberryWeather

# Copy project files to installation directory
cp -r ./* ~/.RaspberryWeather/

# Change to the installation directory
cd ~/.RaspberryWeather

# Build the Go application
echo "Building RaspberryWeather application..."
go build -o RaspberryWeather .

if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

# Make the binary executable
chmod +x RaspberryWeather

# Copy the systemd service file
sudo cp RaspberryWeather.service /etc/systemd/system/RaspberryWeather.service

# Update the service file to use the new path
sudo sed -i "s|/home/pi/RaspberryWeather|$HOME/RaspberryWeather|g" /etc/systemd/system/RaspberryWeather.service

# Reload systemd daemon
sudo systemctl daemon-reload

# Enable and start the service
sudo systemctl enable RaspberryWeather.service
sudo systemctl start RaspberryWeather.service

echo "RaspberryWeather service installed and started successfully!"

# Check service status
sudo systemctl status RaspberryWeather.service