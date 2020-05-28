package test

import (
	"context"
	"os"

	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
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
	cfg, err := config.GetConfigWithContext(context)
	Expect(err).NotTo(HaveOccurred())
	return cfg
}

func MustManager(ctx context.Context, ns string) manager.Manager {
	cfg := MustConfig("")
	return ManagerWithOpts(ctx, cfg, manager.Options{
		Namespace: ns,
		// Disable metrics and health probe to allow tests to run in parallel.
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
}

func ManagerWithOpts(ctx context.Context, cfg *rest.Config, opts manager.Options) manager.Manager {

	mgr, err := manager.New(cfg, opts)
	Expect(err).NotTo(HaveOccurred())

	go func() {
		defer ginkgo.GinkgoRecover()
		err = mgr.Start(ctx.Done())
		Expect(err).NotTo(HaveOccurred())
	}()

	mgr.GetCache().WaitForCacheSync(ctx.Done())

	return mgr
}

func MustManagerNotStarted(ns string) manager.Manager {
	mgr, err := manager.New(MustConfig(""), manager.Options{
		Namespace: ns,
		// Disable metrics and health probe to allow tests to run in parallel.
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
	Expect(err).NotTo(HaveOccurred())
	return mgr
}

func MustApiConfigWithContext(context string) *api.Config {
	clientCfg := ClientConfigWithContext(context)
	cfg, err := clientCfg.ConfigAccess().GetStartingConfig()
	Expect(err).NotTo(HaveOccurred())
	if context != "" {
		cfg.CurrentContext = context
	}
	return cfg
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
