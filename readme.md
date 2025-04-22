# Remote Mouse

A cross-platform solution that allows controlling your computer's mouse remotely using a mobile device.

## Project Structure

This project consists of two main components:

- **remote-mouse-server**: A Go-based server application that runs on your computer and handles mouse control.
- **remote-mouse-app**: A React Native/Expo mobile application that serves as the remote control interface.

## Remote Mouse Server

The server component is written in Go and is responsible for:

- Establishing a connection between your computer and mobile device
- Translating commands from the app into mouse movements and clicks
- Supporting various operating systems (macOS, Windows, Linux)

To learn more about the server, check the [server README](./remote-mouse-server/README.md).

## Remote Mouse App

The mobile app is built with React Native and Expo, providing:

- An intuitive touch interface for controlling your computer's mouse
- Connection management to the server
- Cross-platform support for iOS and Android

To learn more about the app, check the [app README](./remote-mouse-app/README.md).

## Getting Started

1. Start the server on your computer:
   ```
   cd remote-mouse-server
   go run main.go
   ```

2. Launch the mobile app:
   ```
   cd remote-mouse-app
   npm start
   ```

3. Connect your mobile device to the server by entering the server's IP address in the app.

## Development

This project is structured as a monorepo containing both the server and mobile app components. Each has its own development workflow and dependencies.

*This README was enthusiastically crafted by an AI that still can't use a real mouse... oh, the irony!*