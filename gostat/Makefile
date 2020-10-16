all: build

clean:
	@rm -rf gostat

build:
	CGO_ENABLED=0 GOOS=linux go build

run: 
	@go run
