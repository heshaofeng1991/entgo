GO_FILES=$(shell find . -type f -name '*.go')

.PHONY: ent

ent:
	@echo "Generating ent code..."
	@rm -rf ./ent/gen 
	@mkdir -p ./ent/gen
	@cp -r ./ent/internal ./ent/gen
	go generate ./ent
	@cp -r ./ent/gen/internal ./ent

tidy:
	@echo "TIDYING CODE..."
	@go mod tidy -compat=1.19

all: ent tidy