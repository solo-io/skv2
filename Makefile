#----------------------------------------------------------------------------------
# Build
#----------------------------------------------------------------------------------

# Build dependencies

.PHONY: mod-download
mod-download:
	go mod download


DEPSGOBIN=$(shell pwd)/_output/.bin
export GOBIN:=$(DEPSGOBIN)
export PATH:=$(GOBIN):$(PATH)

.PHONY: install-go-tools
install-go-tools: mod-download
	mkdir -p $(DEPSGOBIN)
	go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	go install github.com/solo-io/protoc-gen-ext@v0.0.16
	go install github.com/golang/mock/mockgen@v1.4.4
	go install github.com/onsi/ginkgo/ginkgo@v1.16.4
	go install golang.org/x/tools/cmd/goimports

# proto compiler installation
PROTOC_VERSION:=3.15.8
PROTOC_URL:=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}
.PHONY: install-protoc
ifeq ($(shell $(DEPSGOBIN)/protoc --version | grep -c ${PROTOC_VERSION}),1)
install-protoc:
	@echo expected protoc version ${PROTOC_VERSION} already installed
else
install-protoc:
ifeq ($(shell uname),Darwin)
	@echo downloading protoc for osx
	wget $(PROTOC_URL)-osx-x86_64.zip -O $(DEPSGOBIN)/protoc-${PROTOC_VERSION}.zip
else
ifeq ($(shell uname -m),aarch64)
	@echo downloading protoc for linux aarch64
	wget $(PROTOC_URL)-linux-aarch_64.zip -O $(DEPSGOBIN)/protoc-${PROTOC_VERSION}.zip
else
	@echo downloading protoc for linux x86-64
	wget $(PROTOC_URL)-linux-x86_64.zip -O $(DEPSGOBIN)/protoc-${PROTOC_VERSION}.zip
endif
endif

	unzip $(DEPSGOBIN)/protoc-${PROTOC_VERSION}.zip -d $(DEPSGOBIN)/protoc-${PROTOC_VERSION}
	mv $(DEPSGOBIN)/protoc-${PROTOC_VERSION}/bin/protoc $(DEPSGOBIN)/protoc
	chmod +x $(DEPSGOBIN)/protoc
	rm -rf $(DEPSGOBIN)/protoc-${PROTOC_VERSION} $(DEPSGOBIN)/protoc-${PROTOC_VERSION}.zip
endif

.PHONY: install-tools
install-tools: install-go-tools install-protoc

# Generated Code - Required to update Codgen Templates
.PHONY: generated-code
generated-code: install-tools update-licenses
	$(DEPSGOBIN)/protoc --version
	go run api/generate.go
	# the api/generate.go command is separated out to enable us to run go generate on the generated files (used for mockgen)
# this re-gens test protos
	go test ./codegen
	go generate -v ./...
	$(DEPSGOBIN)/goimports -w .
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
	$(DEPSGOBIN)/goimports -w .

#----------------------------------------------------------------------------------
# Third Party License Management
#----------------------------------------------------------------------------------
.PHONY: update-licenses
update-licenses:
	# check for GPL licenses, if there are any, this will fail
	cd ci/oss_compliance; GO111MODULE=on go run oss_compliance.go osagen -c "GNU General Public License v2.0,GNU General Public License v3.0,GNU Lesser General Public License v2.1,GNU Lesser General Public License v3.0,GNU Affero General Public License v3.0"

	cd ci/oss_compliance; GO111MODULE=on go run oss_compliance.go osagen -s "Mozilla Public License 2.0,GNU General Public License v2.0,GNU General Public License v3.0,GNU Lesser General Public License v2.1,GNU Lesser General Public License v3.0,GNU Affero General Public License v3.0" > osa_provided.md
	cd ci/oss_compliance; GO111MODULE=on go run oss_compliance.go osagen -i "Mozilla Public License 2.0" > osa_included.md

#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

# Important to clean before pushing new releases. Dockerfiles and binaries may not update properly
.PHONY: clean
clean:
	rm -rf codegen/*-packr.go
	rm -rf pkg/api
	rm -rf vendor_any
	rm -rf _output
