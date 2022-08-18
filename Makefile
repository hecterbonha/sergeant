build:
	go build -o bin/sergeant main.go
	
serve:
	./bin/sergeant

dev:
	./bin/air

test:
	go test -v ./...