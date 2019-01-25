GIT_DIR=$(shell git rev-parse --git-dir)

setup_development:
	cp ./scripts/pre-commit.sh $(GIT_DIR)/hooks/pre-commit

build:
	go build -o ./main

test_all:
	rm -rf ./coverage
	mkdir -p ./coverage
	gotest -v -cover -coverprofile ./coverage/cover.out ./...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

syntax_check:
	go vet ./...
	staticcheck ./...
	golint ./...

# clean up go modules
clean_modules:
	go mod tidy
