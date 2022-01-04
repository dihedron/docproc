default: build

.PHONY: build
build:
	goreleaser build --snapshot --single-target --rm-dist

.PHONY: clean
clean:
	@rm -rf dist/ 

