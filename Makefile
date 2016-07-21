deps: 
	govendor update +local
	govendor update +external

start:
	go run api/*.go

test:
	go test ./...

