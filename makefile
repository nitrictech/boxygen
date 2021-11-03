install:
	@echo installing go dependencies
	@go mod download

install-tools: install
	@echo Installing tools from tools.go
	@cat ./tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go get %

generate-proto:
	@echo Generating Proto Sources
	@mkdir -p ./pkg/proto
	@protoc --go_out=./pkg/proto --go-grpc_out=./pkg/proto -I ./proto ./proto/*/**/*.proto

build:
	@CGO_ENABLED=0 go build -tags containers_image_openpgp -o bin/boxygen -ldflags="-extldflags=-static" ./cmd/docker/run.go

build-docker:
	@docker build . -f docker/docker/Dockerfile -t boxygen
	