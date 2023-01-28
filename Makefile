.DEFAULT_GOAL := build

pwd := $(shell pwd)
mod.name := httpstatusdesc
pkg.name := github.com/scottbrown/$(mod.name)
build.dir := $(pwd)/.build

.PHONY: build
build:
	go build -o $(build.dir)/$(mod.name) $(pkg.name)

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf $(build.dir)
