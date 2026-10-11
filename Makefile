.PHONY: build run clean

build:
	go build -o bin/gendiff ./cmd/gendiff

run: build
	./bin/gendiff $(ARGS)

clean:
	rm -rf bin
