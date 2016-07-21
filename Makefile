remove_deps:
	govendor remove +local
	govendor remove +external 

clean_deps_invalid:
	@find vendor -name "*_test.go" -delete

_deps:
	govendor update +local
	govendor update +external 

deps: remove_deps _deps clean_deps_invalid

start: deps
	go run api/*.go

test: clean_deps_invalid
	go test ./...

