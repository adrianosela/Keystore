NAME:=$(shell basename `git rev-parse --show-toplevel`)
HASH:=$(shell git rev-parse --verify --short HEAD)

all: build

clean:
	rm -rf pkg bin

build:
	GOOS=linux GOARCH=amd64 go build -a -o $(NAME)
	docker build -t keystore .
