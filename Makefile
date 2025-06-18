.PHONY: all

all: clean build

clean:
	git add $$(gofmt -w ./..)
	go mod tidy
	git add go.mod go.sum

build:
	go build -o ./bin/khata github.com/sreekar2307/khata/cmd

http: 
	./bin/khata http

token: 
	./bin/khata token $(user) $(password)

seed: 
	./bin/khata seed

migrate: 
	./bin/khata migrate
