build-bin:
	go build -ldflags "-s -w" -o ./build/address-suggestion-proxy ./main.go

build-docker:
	docker build -t address-suggestion-proxy:local .

build-clean:
	rm -rf ./build

run-docker-dev:
	docker-compose -f docker-compose.dev.yml up

run-docker-prod:
	docker-compose -f docker-compose.dev.yml up -d

run:
	docker run --rm -it -v $(shell pwd):/app/ --name address-suggestion-proxy --entrypoint 'sh' address-suggestion-proxy:local

test:
	go test -v ./...

lint:
	golangci-lint run --out-format=github-actions

mod-update:
	go get -u && go mod tidy