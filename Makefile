COVERAGE_MIN ?= 80
COVERAGE_PROFILE := coverage.out
# Package main is excluded from coverage.
COVERAGE_PKGS := $(shell go list ./... | grep -v /cmd/)

.PHONY: build run test test-coverage lint clean

build:
	go build -o bin/gendiff ./cmd/gendiff

run: build
	./bin/gendiff $(ARGS)

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=$(COVERAGE_PROFILE) $(COVERAGE_PKGS)
	@total=$$(go tool cover -func=$(COVERAGE_PROFILE) | awk '/^total:/ { sub("%", "", $$3); print $$3 }'); \
	echo "Total coverage: $$total% (min $(COVERAGE_MIN)%)"; \
	awk -v total="$$total" -v min="$(COVERAGE_MIN)" 'BEGIN { exit !(total >= min) }' || \
		{ echo "Coverage is below $(COVERAGE_MIN)%"; exit 1; }

lint:
	golangci-lint run ./...

clean:
	rm -rf bin $(COVERAGE_PROFILE)
