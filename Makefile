build:
	go build -o bin/sergeant -tags=go_json
	
serve:
	./bin/sergeant

dev:
	./bin/air

ui:
	cd web && pnpm run dev

test:
	go test -v ./...