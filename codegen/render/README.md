### Testing

In order to run all tests in this package you must set the REMOTE_CLUSTER_CONTEXT env var to the name of a context in
your kubeconfig. To quickly provision two [kind](https://github.com/kubernetes-sigs/kind) clusters for testing, run
`./ci/setup-kind.sh <remote-context-name>`. This will create two kind clusters, `skv2-test-master` and
`<remote-context-name>`, and set your current context to `kind-skv2-test-master`. To run the multicluster tests in this
package, you can then use `REMOTE_CLUSTER_CONTEXT=kind-<remote-context-name> TEST_PKG=codegen/render make run-tests`.
To cleanup the kind clusters, run `./ci/teardown-kind.sh`.