FROM golang as build

WORKDIR /app

RUN apt-get update && apt-get install -y protobuf-compiler

COPY go.mod go.sum ./
COPY makefile makefile
COPY tools/tools.go tools/tools.go
RUN make install-tools

COPY . .

RUN make build

FROM alpine

COPY --from=build /app/bin/boxygen /boxygen
RUN chmod +rx /boxygen

# Document default gRPC port
EXPOSE 50051

# Document workspace volume for building
VOLUME "/workspace/"

WORKDIR /

ENTRYPOINT [ "/boxygen" ]

