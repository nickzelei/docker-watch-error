
all: build

build:
	GOOS=linux go build -o bin/api main.go
