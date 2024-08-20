LDFLAGS += -X 'main.ReleaseVersion=$(shell git describe --tags || echo "development")'
LDFLAGS += -X 'main.GitHash=$(shell git rev-parse HEAD)'

.PHONY: all
all: test

test:
	go test -v
