# ledserver

A Go HTTP server running on a Raspberry Pi 5 that controls an LED wired to an Arduino Uno over serial.

## What it does

- Register, login, logout with bcrypt password hashing
- Bearer token authentication
- Protected routes via middleware
- POST /led/on and POST /led/off send commands to the Arduino over serial
- Arduino controls a physical LED on a breadboard
- Minimal frontend — login form, two buttons
- Runs as a systemd service, starts on boot
- udev rule gives the Arduino a stable device name (/dev/arduino)

## What was new

- Writing a systemd service on Linux
- Cross compiling Go for ARM64
- Serial communication between Pi and Arduino
- udev rules
- Arduino sketch in C
- Middleware and bearer token auth in Gin
- Basic frontend with fetch API

## Stack

- Go + Gin
- SQLite via modernc.org/sqlite
- bcrypt via golang.org/x/crypto
- go.bug.st/serial
- Vanilla HTML + JavaScript
- Raspberry Pi 5 + Arduino Uno R4
- Tailscale for remote access

## Routes

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | / | No | Frontend |
| POST | /register | No | Create account |
| POST | /login | No | Get session token |
| POST | /logout | Yes | Delete session |
| POST | /led/on | Yes | Turn LED on |
| POST | /led/off | Yes | Turn LED off |

## What's next

More electronics — motors, sensors, lower level boards. Better security. More complex projects.