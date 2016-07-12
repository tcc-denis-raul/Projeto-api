remove_deps:
	@rm -rf $(GOPATH)/src/github.com/tcc-denis-raul

deps: remove_deps
	go get -d ./...

start: deps
	./run.sh

test:
	go test ./...

