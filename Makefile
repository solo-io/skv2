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

# Generated Code - Required to update Codgen Templates
.PHONY: generated-code
generated-code:
	go generate ./...

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
		-skipPackage=$(SKIP_PACKAGES) $(TEST_PKG)


#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

# Important to clean before pushing new releases. Dockerfiles and binaries may not update properly
.PHONY: clean
clean:
	rm -rf codegen/*-packr.go