## Getting Started

***PREREQUISITE: [Go is installed](https://go.dev/doc/install) and [protoc is installed](https://grpc.io/docs/languages/go/quickstart/)***
***Add the go env path to ~/.bashrc or ~/.zshrc depending on the shell and source it in order to implement terminal commands properly***

if unix-like(linux):

```bash
vim ~/.bashrc
```
or if MacOS:

```bash
vim ~/.zshrc
```

to add the following line

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

and source it

```bash
source ~/.bashrc
```

or

```bash
source ~/.zshrc
```

***ENVIRONMENT: tested and run on Linux and MacOS***

The below instructions are how the project was configured to be working from the scratch

## Running Each Sample Project

1. Create proto code "[PROTOBUF_NAME].proto" for service and message

2. Create a folder for package(pkg) in which protoc generated codes will belong to

```bash
mkdir [PKG_NAME]
```

3. Compile "[PROTOBUF_NAME].proto" with protoc

```bash
protoc --go_out=[PKG_NAME] --go_opt=paths=source_relative --go-grpc_out=[PKG_NAME] --go-grpc_opt=paths=source_relative [PROTOBUF_NAME].proto
```

4. Initialize go module with the package that is of the same name as protoc generated package; this is to implement correct go packages import path

```bash
go mod init [PKG_NAME]
```

In order to make the module file tidy:

```bash
go mod tidy
```

5. Get gRPC package

```bash
go get -u google.golang.org/grpc
```

6. Write server and client: server.go and client.go

7. Run each server and client in sequence in separate terminal

in a terminal:

```bash
go run server.go
```

in another terminal:

```bash
go run client.go
```