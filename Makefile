#----------------------------------------------------------------------------------
# Build
#----------------------------------------------------------------------------------

# Build dependencies

.PHONY: mod-download
mod-download:
	go mod download


DEPSGOBIN=$(shell pwd)/_output/.bin

.PHONY: install-go-tools
install-go-tools: mod-download
	mkdir -p $(DEPSGOBIN)
	GOBIN=$(DEPSGOBIN) go install github.com/gobuffalo/packr/packr
	GOBIN=$(DEPSGOBIN) go install github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(DEPSGOBIN) go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	GOBIN=$(DEPSGOBIN) go install github.com/solo-io/protoc-gen-ext
	GOBIN=$(DEPSGOBIN) go install github.com/golang/mock/mockgen
	GOBIN=$(DEPSGOBIN) go install golang.org/x/tools/cmd/goimports

# Generated Code - Required to update Codgen Templates
.PHONY: generated-code
generated-code: clean install-go-tools
	PATH=$(DEPSGOBIN):$$PATH go run api/generate.go
	# the api/generate.go command is separated out to enable us to run go generate on the generated files (used for mockgen)
	PATH=$(DEPSGOBIN):$$PATH go generate -v ./...
	PATH=$(DEPSGOBIN):$$PATH goimports -w .
	go mod tidy

#----------------------------------------------------------------------------------
# Test
#----------------------------------------------------------------------------------

# run all tests
# set TEST_PKG to run a specific test package
.PHONY: run-tests
run-tests:
	PATH=$(DEPSGOBIN):$$PATH ginkgo -r -failFast -trace -progress \
		-progress \
		-compilers=4 \
		-skipPackage=$(SKIP_PACKAGES) $(TEST_PKG) \
		-failOnPending \
		-randomizeAllSpecs \
		-randomizeSuites \
		-keepGoing
	goimports -w .

#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

# Important to clean before pushing new releases. Dockerfiles and binaries may not update properly
.PHONY: clean
clean:
	rm -rf codegen/*-packr.go
	rm -rf pkg/api
	rm -rf vendor_any
