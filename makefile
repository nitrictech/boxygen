install:
	@echo installing go dependencies
	@go mod download

install-tools: install
	@echo Installing tools from tools.go
	@cat ./tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

test:
	@echo Running Tests
	@go run github.com/onsi/ginkgo/ginkgo ./pkg/...

generate-proto:
	@echo Generating Proto Sources
	@mkdir -p ./pkg/proto
	@protoc --go_out=./pkg/proto --go-grpc_out=./pkg/proto -I ./proto ./proto/*/**/*.proto

build: generate-proto
	@CGO_ENABLED=0 go build -tags containers_image_openpgp -o bin/boxygen -ldflags="-extldflags=-static" ./cmd/docker/run.go

build-docker:
	@DOCKER_BUILDKIT=1 docker build . -f docker/docker/Dockerfile -t nitrictech/boxygen

sourcefiles := $(shell find . -type f -name "*.go")

license-header-check:
	@echo Checking License Headers to Source Files
	@addlicense -check -c "Nitric Technologies Pty Ltd." -y "2021" $(sourcefiles)

license-check: install-tools license-check-dev license-check-aws license-check-gcp license-check-azure
	@echo Checking OSS Licenses