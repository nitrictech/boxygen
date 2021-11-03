FROM golang:alpine as gobuild

WORKDIR /app

COPY . .
# Build static binary for use in final container
RUN CGO_ENABLED=0 go build -tags containers_image_openpgp -o bin/boxygen -ldflags="-extldflags=-static" ./cmd/docker/run.go

FROM docker:dind

COPY --from=gobuild /app/bin/boxygen /boxygen
RUN chmod +rx /boxygen

WORKDIR /

ENTRYPOINT [ "/boxygen" ]

