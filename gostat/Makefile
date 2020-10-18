all: build

clean:
	@rm -rf gostat

build:
	CGO_ENABLED=0 GOOS=linux go build
	@cp config.json.example config.json

run: 
	@go run
