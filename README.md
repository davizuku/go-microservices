# go-microservices

Test code following the set of videos from:  https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_

Teacher's repository: https://github.com/PacktPublishing/Building-Microservices-with-Go-Second-Edition

## Start working

Start docker running go image:

```
make run
```

_Enter_ docker using:

```
make bash
```

### Start gRPC server

```
make bash
go run grpc/main.go
```

### gRPC curl

- Install the command `grpcurl` with:

```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

- List available services in the GRPC server:

```
grpcurl --plaintext localhost:3001 list
```

- Describe `Currency` server:

```
grpcurl --plaintext localhost:3001 describe Currency
```

- You can also use `describe` with other elements of the service:

```
grpcurl --plaintext localhost:3001 describe .RateRequest
```

- Send request with data:
```
grpcurl --plaintext  -d '{"Base": "USD", "Destination": "EUR"}' localhost:3001 Currency.GetRate
```

## Finish working

Stop any running docker:

```
make down
```

## Useful links:

- https://golang.org/pkg/
- http://www.gorillatoolkit.org/pkg/mux
- https://github.com/golang/go/wiki/Modules
- https://github.com/go-playground/validator
- https://goswagger.io
- https://github.com/golang-standards/project-layout
- https://redocly.github.io/redoc/

### gRPC
- https://grpc.io/docs/quickstart/go/
- https://developers.google.com/protocol-buffers/docs/proto3
- https://github.com/fullstorydev/grpcurl


### Debugging Go in VS Code

- https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code
- https://github.com/go-delve/delve/issues/986
- make sure you execute `xcode-select --install`
