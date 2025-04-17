BIN_NAME={{ cookiecutter.bin_name }}
VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "-DIRTY" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

default: fmt lint modernize test build

help:
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make test            Run tests.'
	@echo '    make clean           Clean the directory tree.'
	@echo '    make upgrade         Upgrades all dependencies.'
	@echo '    make modernize       Find and fix more modern go usages.'
	@echo '    make fmt             Format sources.'
	@echo '    make lint            Run linters.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X {{ cookiecutter.module_name }}/version.gitRef=${GIT_COMMIT}${GIT_DIRTY} -X {{ cookiecutter.module_name }}/version.buildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

fmt:
	go fmt ./...

upgrade:
	go get -u ./...

lint:
	go vet ./...
	staticcheck ./...
	golangci-lint run

modernize:
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest  ./...
