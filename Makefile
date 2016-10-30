remove_deps:
	govendor remove +v

_deps:
	go get ./...
	govendor add +external

deps: remove_deps _deps

run_:
	build/projeto-api

run: deps remove_build build run_

test:
	go test ./...

build:
	go build -o build/projeto-api api/*

remove_build:
	@rm -r build 2> /dev/null; true
