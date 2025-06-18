.PHONY: all

all: clean build

clean:
	git add $$(gofmt -w ./..)
	go mod tidy
	git add go.mod go.sum

build:
	go build -o ./bin/khata github.com/sreekar2307/khata/cmd

http: build
	./bin/khata http

token: build
	./bin/khata token $(user) $(password)

seed: build
	./bin/khata seed

migrate: build
	./bin/khata migrate
