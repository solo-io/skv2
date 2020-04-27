#!/bin/bash

set -ex

if [ "$1" == "cleanup" ]; then
  kind delete cluster --name skv2-test-master
  kind delete cluster --name skv2-test-remote
  exit 0
fi

kind create cluster --name skv2-test-master
kind create cluster --name skv2-test-remote

kubectl config use-context kind-skv2-test-master