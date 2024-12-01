BIN_NAME={{ cookiecutter.bin_name }}
VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "-DIRTY" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

default: test

help:
	@echo 'Management commands for {{ cookiecutter.app_name }}:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make clean           Clean the directory tree.'
	@echo '    make upgrade         Upgrades all dependencies.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X {{ cookiecutter.module_name }}/version.gitRef=${GIT_COMMIT}${GIT_DIRTY} -X {{ cookiecutter.module_name }}/version.buildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...


upgrade:
	go get -u ./...
