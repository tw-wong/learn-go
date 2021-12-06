# Protocol Buffers

## Make sure Protobuf compiler has installed:
```sh
~ which protoc
/opt/homebrew/bin/protoc
~ protoc --version
libprotoc 3.17.3
```

## Download dependencies
```sh
~ go mod download
~ go mod tidy
```

# Compile Protobuf
```sh
~ protoc --go_out=. ./protobuf/*.proto
```
- Generate *.pb.go files at "gen" directory.

## Run
```sh
~ go run main.go
Marshal data: [10 5 65 108 105 99 101 16 24 24 1]
Print Person:
- Age: 24
- Name: Alice
- Gender: FEMALE
```
```