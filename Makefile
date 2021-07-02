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

# Dependencies for code generation
.PHONY: mockgen protoc-gen-go protoc-gen-ext protoc-gen-jsonshim protoc-plugins

mockgen: $(DEPSGOBIN)/mockgen
$(DEPSGOBIN)/mockgen:
	go build -o $@ github.com/golang/mock/mockgen

protoc-gen-go: $(DEPSGOBIN)/protoc-gen-go
$(DEPSGOBIN)/protoc-gen-go:
	go build -o $@ github.com/golang/protobuf/protoc-gen-go

protoc-gen-ext: $(DEPSGOBIN)/protoc-gen-ext
$(DEPSGOBIN)/protoc-gen-ext:
	go build -o $@ github.com/solo-io/protoc-gen-ext

protoc-plugins: protoc-gen-go protoc-gen-ext # protoc-gen-jsonshim

# Generated Code - Required to update Codgen Templates
.PHONY: generated-code
generated-code: update-licenses mockgen protoc-plugins
	go run api/generate.go
	# the api/generate.go command is separated out to enable us to run go generate on the generated files (used for mockgen)
# this re-gens test protos
	go test ./codegen
	go generate -v ./...
	go run golang.org/x/tools/cmd/goimports -w .
	go mod tidy

#----------------------------------------------------------------------------------
# Test
#----------------------------------------------------------------------------------

# run all tests
# set TEST_PKG to run a specific test package
.PHONY: run-tests
run-tests: protoc-plugins
	go run github.com/onsi/ginkgo/ginkgo -r -failFast -trace -progress \
		-progress \
		-compilers=4 \
		-skipPackage=$(SKIP_PACKAGES) $(TEST_PKG) \
		-failOnPending \
		-randomizeAllSpecs \
		-randomizeSuites \
		-keepGoing
	go run golang.org/x/tools/cmd/goimports -w .

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
