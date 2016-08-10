remove_deps:
	govendor remove +local
	govendor remove +external 

clean_deps_invalid:
	@find vendor -name "*_test.go" -delete

_deps:
	govendor update +local
	govendor update +external 

deps: remove_deps _deps clean_deps_invalid

run_:
	build/projeto-api

run: deps remove_build build run_

test: clean_deps_invalid
	go test ./...

build:
	go build -o build/projeto-api api/*

remove_build:
	@rm -r build 2> /dev/null; true
