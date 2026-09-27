# ledserver

A Go HTTP server running on a Raspberry Pi 5 that controls an LED wired to an Arduino Uno over serial.

Built with Gin, SQLite, bcrypt, and a minimal HTML frontend. No tutorials were followed end to end — this was built piece by piece, with AI assistance along the way, but the learning was real.

## What it does

- Register, login, logout with bcrypt password hashing
- Bearer token authentication
- Protected routes via middleware — unauthenticated requests are blocked
- POST /led/on and POST /led/off send serial commands to the Arduino
- Arduino controls a physical LED on a breadboard
- Minimal frontend — login form, two buttons, served by the Go server itself
- Runs as a systemd service on the Pi, starts automatically on boot
- udev rule gives the Arduino a stable device name (/dev/arduino)

## What was new territory

- Writing a systemd service on Linux
- Cross compiling Go for ARM64
- Serial communication between Pi and Arduino
- udev rules for stable device naming
- Arduino sketch in C
- Bearer token auth and middleware in Gin
- A little bit of frontend (HTML, fetch API)

## What's next

- Learn more about electronics — motors, sensors, more complex automation
- Explore lower level boards (ESP32, bare microcontrollers)
- Security hardening — this server is not production ready
- More complex project structures and database design
- Python backend (FastAPI or Flask) for comparison

## Stack

- Go + Gin
- SQLite via modernc.org/sqlite
- bcrypt via golang.org/x/crypto
- go.bug.st/serial for Arduino communication
- Vanilla HTML + JavaScript frontend
- Raspberry Pi 5 + Arduino Uno R4
- Tailscale for remote access

## Routes

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | / | No | Serves the frontend |
| POST | /register | No | Create account |
| POST | /login | No | Get session token |
| POST | /logout | Yes | Delete session |
| POST | /led/on | Yes | Turn LED on |
| POST | /led/off | Yes | Turn LED off |

## The moment that made it worth it

Clicking a button on my phone and watching a physical LED turn on instantly. Not on a simulator. Not on my laptop. On a breadboard, wired to an Arduino, controlled by a server running on a Pi across the room.

That's what made all the debugging worth it.