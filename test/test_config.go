package test

import (
	"context"
	"os"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	clientconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	// visible for testing
	ServiceAccountClusterAdminRoles = []client.ObjectKey{
		{
			Name: "cluster-admin",
		},
	}
)

func MustConfig(context string) *rest.Config {
	cfg, err := clientconfig.GetConfigWithContext(context)
	Expect(err).NotTo(HaveOccurred())
	return cfg
}

func MustManager(ctx context.Context, ns string) manager.Manager {
	cfg := MustConfig("")
	return ManagerWithOpts(ctx, cfg, manager.Options{
		Cache: cache.Options{
			DefaultNamespaces: map[string]cache.Config{ns: cache.Config{}},
		},
		Controller: config.Controller{
			// see https://github.com/kubernetes-sigs/controller-runtime/issues/2937
			// in short, our tests reuse the same name (reasonably so) and the controller-runtime
			// package does not reset the stack of controller names between tests, so we disable
			// the name validation here.
			SkipNameValidation: ptr.To(true),
		},
		// Disable metrics and health probe to allow tests to run in parallel.
		Metrics: server.Options{
			BindAddress: "0",
		},
		HealthProbeBindAddress: "0",
	})
}

func ManagerWithOpts(ctx context.Context, cfg *rest.Config, opts manager.Options) manager.Manager {

	mgr, err := manager.New(cfg, opts)
	Expect(err).NotTo(HaveOccurred())

	go func() {
		defer ginkgo.GinkgoRecover()
		err = mgr.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	}()

	mgr.GetCache().WaitForCacheSync(ctx)

	return mgr
}

func ClientConfigWithContext(context string) clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = os.Getenv("KUBECONFIG")
	// NOTE: ConfigOverrides are NOT propagated to `GetStartingConfig()`, so we set CurrentContext on the resulting config
	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{
		CurrentContext: context,
	})
	return cfg
}
