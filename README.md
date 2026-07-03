# ByteBridge

ByteBridge is a small Go prototype for sending files over TCP on a local network.

It currently includes:

- A server that listens for incoming TCP connections.
- A client that sends one file to the server.
- A simple protocol: one JSON header, then the raw file bytes.

## Requirements

- Go `1.26.2` or a compatible version.
- `make` for the included development commands.

## Usage

Create the folder used by the server:

```bash
mkdir -p files/server
```

Start the server:

```bash
make server
```

By default, the server listens on port `4000`.

In another terminal, send a file:

```bash
make client file=/path/to/file
```

Received files are saved as:

```text
files/server/received_<filename>
```

## Commands

```bash
make fmt
make vet
make server
make client file=/path/to/file
make buildserver
make buildclient
```

Build outputs are written to:

```text
bin/server
bin/client
```

## Project Layout

```text
cmd/client/       Client entry point
cmd/server/       Server entry point
internals/        TCP, job, and connection code
files/            Local file storage
config.yml        Planned configuration
Makefile          Development commands
```

## Status

This project is still in active development. The basic single-file transfer flow works.

Planned work includes multiple file transfers, device discovery, checksum validation, progress output, and fuller use of `config.yml`.
