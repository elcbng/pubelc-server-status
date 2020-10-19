all: build

clean:
	@rm -rf gostat-*

build:
	CGO_ENABLED=0 GOOS=linux go build -o gostat-$(shell uname -m)-native -ldflags "-s -w"
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o gostat-mipsle -ldflags "-s -w"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o gostat-armv7 -ldflags "-s -w"
	@cp config.json.example config.json

run: 
	@go run
