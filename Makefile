build:
	go build -o bin/sergeant main.go
	
serve:
	./bin/sergeant

dev:
	bash dev.sh
 
test:
	go test -v ./...