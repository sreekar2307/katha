.PHONY: all

all: clean build

clean:
	git add $$(gofmt -w ./..)
	go mod tidy
	git add go.mod go.sum

build:
	go build -o ./bin/khata github.com/sreekar2307/khata/cmd

run_http: build
	./bin/katha http

token: build
	./bin/katha token $(user) $(password)

seed: build
	./bin/katha seed

migrate: build
	./bin/katha migrate
