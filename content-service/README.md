# Content Service gRPC API (Go + MongoDB)

This project is a gRPC service in Go that exposes two methods: `GetContentById` and `ListContent`. It uses MongoDB as the backend and Protocol Buffers for API definition.

---

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
  - [1. Install System Dependencies](#1-install-system-dependencies)
  - [2. Clone and Prepare the Project](#2-clone-and-prepare-the-project)
  - [3. Install Go Dependencies](#3-install-go-dependencies)
  - [4. Proto Generation](#4-proto-generation)
  - [5. Environment Variables](#5-environment-variables)
  - [6. Running the Server](#6-running-the-server)
- [Development Notes](#development-notes)
- [Troubleshooting](#troubleshooting)

---

## Features
- gRPC API with two methods: `GetContentById` and `ListContent`
- MongoDB integration
- Protobuf-based schema
- Environment variable support via `.env`

## Project Structure
```
content-service/
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── src/
    ├── main.go
    ├── server.go
    ├── db/
    │   └── db.go
    └── protos/
        ├── content.pb.go
        └── content_grpc.pb.go
```

## Setup Instructions

### 1. Install System Dependencies
- **Go** (>= 1.18): https://go.dev/doc/install
- **protoc** (Protocol Buffers compiler):
  ```zsh
  brew install protobuf
  ```
- **protoc Go plugins**:
  ```zsh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
  Ensure `$HOME/go/bin` is in your `PATH`:
  ```zsh
  echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.zshrc
  source ~/.zshrc
  ```

### 2. Clone and Prepare the Project
Clone this repo and `cd` into the project directory.

### 3. Install Go Dependencies
```zsh
cd content-service
# Install all Go dependencies
go mod tidy
```

### 4. Proto Generation
- Make sure your proto file is at `../protos/content-service/content.proto`.
- Generate Go code from proto:
  ```zsh
  make proto
  ```
  This will generate `content.pb.go` and `content_grpc.pb.go` in `src/protos/`.

### 5. Environment Variables
- Create a `.env` file in the project root:
  ```env
  DATABASE_URL=mongodb+srv://<user>:<password>@<host>/<db>?retryWrites=true&w=majority
  ```
- The service loads this automatically at startup.

### 6. Running the Server
```zsh
go run ./src
```
The server listens on `:50051` by default.

---

## Development Notes
- The gRPC service is defined in the proto file and generated into Go code.
- MongoDB connection is managed as a singleton in `db/db.go`.
- The server implements the gRPC methods in `server.go`.
- Timestamps are returned as `google.protobuf.Timestamp` objects (with `seconds` and `nanos`).

## Troubleshooting

### Common Errors & Fixes
- **protoc-gen-go: program not found or is not executable**
  - Ensure you ran the `go install ...` commands above and `$HOME/go/bin` is in your `PATH`.
- **protoc-gen-go: unable to determine Go import path for ...**
  - Add a `go_package` option to your proto file.
- **Response message parsing error: invalid wire type ...**
  - Ensure your MongoDB `_id` is mapped to a string in the response, and that all fields match the proto types.
- **Timestamps appear as `{seconds: ..., nanos: ...}`**
  - This is the correct protobuf format. Convert to a string on the client if needed.

---

## Credits
- Go gRPC: https://grpc.io/docs/languages/go/
- MongoDB Go Driver: https://pkg.go.dev/go.mongodb.org/mongo-driver
- Protocol Buffers: https://developers.google.com/protocol-buffers
