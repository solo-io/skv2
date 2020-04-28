#!/bin/bash

set -ex

#Delete all kind clusters
kind get clusters | while read -r r; do kind delete cluster --name "$r"; done
exit 0