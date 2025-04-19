# Learn gRPC

A simple gRPC demo project implementing a user service with unary and server streaming endpoints.

## Project Structure

```
├── client/         # gRPC client implementation
├── learn/          # Proto generated code
├── server/         # gRPC server implementation
├── Makefile        # Build commands
└── go.mod         # Go module file
```

## Features

- Unary RPC: `GetUser` - Fetch a single user by ID
- Server Streaming RPC: `GetUsers` - Stream a list of users

## Prerequisites

- Go 1.23.5 or later
- Protocol Buffers compiler (`protoc`)
- Go gRPC plugin
- Go Protobuf plugin

## Installation

1. Clone the repository
2. Install dependencies:
```sh
go mod download
```

## Generate gRPC Code

Use the Makefile to generate gRPC code from proto files:

```sh
make gen
```

## Running the Application

1. Start the server:
```sh
go run server/main.go
```

2. Run the client (in a new terminal):
```sh
go run client/main.go --userId=1
```

## Service Definition

The service is defined in `learn/learn.proto`:

- `GetUser`: Returns a single user for a given ID
- `GetUsers`: Streams a list of predefined users

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request