#----------------------------------------------------------------------------------
# Build
#----------------------------------------------------------------------------------

# Build dependencies

.PHONY: mod-download
mod-download:
	go mod download

.PHONY: install-deps
install-deps: mod-download
	go get -v github.com/gobuffalo/packr/packr
	go get -v istio.io/tools/cmd/protoc-gen-jsonshim
	go get -v github.com/gogo/protobuf/protoc-gen-gogo
	go get -v github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	go get -v github.com/solo-io/protoc-gen-ext
	go get -v github.com/golang/mock/mockgen
	go get -v golang.org/x/tools/cmd/goimports@v0.0.0-20200427205912-352a5409fae0
	go mod tidy

# Generated Code - Required to update Codgen Templates
.PHONY: generated-code
generated-code: install-deps
	go generate ./...
	goimports -w .

#----------------------------------------------------------------------------------
# Test
#----------------------------------------------------------------------------------

# run all tests
# set TEST_PKG to run a specific test package
.PHONY: run-tests
run-tests:
	ginkgo -r -failFast -trace -progress \
		-progress \
		-compilers=4 \
		-skipPackage=$(SKIP_PACKAGES) $(TEST_PKG) \
		-failOnPending \
		-randomizeAllSpecs \
		-randomizeSuites
	goimports -w .

#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

# Important to clean before pushing new releases. Dockerfiles and binaries may not update properly
.PHONY: clean
clean:
	rm -rf codegen/*-packr.go