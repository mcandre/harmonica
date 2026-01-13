# INSTALL

We support several installation methods.

# PRECOMPILED BINARIES

https://github.com/mcandre/harmonica/releases

## Requirements

(None)

## Instructions

1. Download release archive.
2. Extract archive.
3. Select executables for your target platform.
4. Copy executabless to a convenient location, e.g. `$HOME/bin`.
5. Ensure location is registered in `$PATH`.

# DOCKER

## Requirements

* [Docker](https://www.docker.com/) 28.0.1+

## Instructions

```sh
docker pull n4jm4/harmonica
```

# BUILD FROM SOURCE

## Requirements

* [Go](https://go.dev/) 1.25.4+
* Ensure `GOBIN` is registered in `$PATH`. Validate like `go env GOBIN; echo "$PATH"`

## Instructions

```sh
go get -tool github.com/mcandre/harmonica/src/cmd/harmonica
go mod tidy
go mod vendor
go install tool
```

For more information on developing harmonica itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
