# Mindera Weekly GRPC

### Installation

- Install Protoc compiler
- Install Proto generator for desired language (C++, java, Kotlin, Python, Go, Ruby, Objective-C, C#, PHP)

#### Compiler

https://grpc.io/docs/protoc-installation/

linux:

```shell
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

Mac:

```shell
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

#### CodeGen

https://developers.google.com/protocol-buffers/docs/reference/overview

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Generate Protoc

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```


## Usefull Links
- https://github.com/protocolbuffers/protobuf-go
- https://developers.google.com/protocol-buffers/docs/proto3