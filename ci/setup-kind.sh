#!/bin/bash

set -ex

if [ "$1" == "" ]; then
  echo "please provide a name for the 'remote' cluster"
  exit 0
fi

kind create cluster --name skv2-test-master
kind create cluster --name "$1"

kubectl config use-context kind-skv2-test-master