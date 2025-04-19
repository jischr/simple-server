# Simple Worker

A simple worker application built with [Echo](https://echo.labstack.com/), a high-performance, extensible, and minimalist Go web framework.

## Features

- Exposes a `/version` endpoint to display the current version of the worker.
- Exposes a `/status` endpoint to check if the worker is running.
- Graceful shutdown on receiving interrupt signals.

## Usage

To run:
```sh
go run -ldflags "-X main.port=8080 -X main.version=1.0.0" main.go
```

To build:
```sh
go build -ldflags "-X main.version=<version> -X main.port=<port>"
```

Replace `<version>` with the desired version string and `<port>` with the desired port number.
