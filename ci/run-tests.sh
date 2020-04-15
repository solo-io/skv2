#!/bin/bash

set -ex

go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

make run-tests