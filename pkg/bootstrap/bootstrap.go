package bootstrap

import (
	"context"
	"net/http"

	"github.com/rotisserie/eris"

	"github.com/solo-io/go-utils/errgroup"

	"github.com/solo-io/skv2/pkg/stats"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/go-logr/zapr"
	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster/watch"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log"
	zaputil "sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// required import to enable kube client-go auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// StartParameters specify paramters for starting a generic controller which may need access to its local cluster as well as remote (multicluster) clients and managers
type StartParameters struct {
	MasterManager   manager.Manager
	McClient        multicluster.Client    // nil if running in agent mode
	Clusters        multicluster.Interface // nil if running in agent mode
	SnapshotHistory *stats.SnapshotHistory
	// Reference to Settings object this controller uses.
	SettingsRef v1.ObjectRef
	// enable additional logging
	VerboseMode bool
}

// the start function that will be called with the initialized parameters
type StartFunc func(
	ctx context.Context,
	parameters StartParameters,
) error

// bootstrap options for starting discovery
type Options struct {
	// MetricsBindPort is the TCP port that the controller should bind to
	// for serving prometheus metrics.
	// It can be set to 0 to disable the metrics serving.
	MetricsBindPort uint32

	// MasterNamespace if specified restricts the Master manager's cache to watch objects in the desired namespace.
	// Defaults to all namespaces.
	//
	// Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace.
	MasterNamespace string

	// enables verbose mode
	VerboseMode bool

	// enables json logger (instead of table logger)
	// NOTE: DO NOT set this to false in Prod, it will crash on DPanic
	JSONLogger bool

	// ManagementContext if specified read the KubeConfig for the management cluster from this context. Only applies when running out of cluster.
	ManagementContext string

	// Reference to the Settings object that the controller should use.
	SettingsRef v1.ObjectRef
}

// convenience function for setting these options via spf13 flags
func (opts *Options) AddToFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&opts.MasterNamespace, "namespace", "n", metav1.NamespaceAll, "if specified restricts the master manager's cache to watch objects in the desired namespace.")
	flags.Uint32Var(&opts.MetricsBindPort, "metrics-port", opts.MetricsBindPort, "port on which to serve Prometheus metrics. set to 0 to disable")
	flags.BoolVar(&opts.VerboseMode, "verbose", true, "enables verbose/debug logging")
	flags.StringVar(&opts.ManagementContext, "context", "", "If specified, use this context from the selected KubeConfig to connect to the local (management) cluster.")
	flags.StringVar(&opts.SettingsRef.Name, "settings-name", opts.SettingsRef.Name, "The name of the Settings object this controller should use.")
	flags.StringVar(&opts.SettingsRef.Namespace, "settings-namespace", opts.SettingsRef.Namespace, "The namespace of the Settings object this controller should use.")

	// This flag disables prod mode when set to false, in other words setting debug to true,
	// Which will cause the app to panic on DPanic.
	flags.BoolVar(&opts.JSONLogger, "prod-logging-mode", true, "Default: true. Set this value to false to enable debug panic logs for development.")
	flags.MarkHidden("prod-logging-mode")
}

// Start a controller with the given start func. The StartFunc will be called with a bootstrapped local manager. If localMode is false, the StartParameters will include initialized multicluster components.
func Start(ctx context.Context, start StartFunc, opts Options, schemes runtime.SchemeBuilder, localMode bool) error {
	return StartMulti(ctx, map[string]StartFunc{"": start}, opts, schemes, localMode)
}

// Like Start, but runs multiple StartFuncs concurrently
func StartMulti(ctx context.Context, startFuncs map[string]StartFunc, opts Options, schemes runtime.SchemeBuilder, localMode bool, addStatsHandlers ...func(mux *http.ServeMux, profiles map[string]string)) error {
	setupLogging(opts.VerboseMode, opts.JSONLogger)

	mgr, err := makeMasterManager(opts, schemes)
	if err != nil {
		return err
	}

	snapshotHistory := stats.NewSnapshotHistory()

	stats.MustStartServerBackground(snapshotHistory, opts.MetricsBindPort, addStatsHandlers...)

	var (
		clusterWatcher multicluster.Interface
		mcClient       multicluster.Client
	)

	if !localMode {
		// construct multicluster watcher and client
		clusterWatcher = watch.NewClusterWatcher(ctx, manager.Options{
			Namespace: "", // TODO (ilackarms): support configuring specific watch namespaces on remote clusters
			Scheme:    mgr.GetScheme(),
		})

		mcClient = multicluster.NewClient(clusterWatcher)
	}

	params := StartParameters{
		MasterManager:   mgr,
		McClient:        mcClient,
		Clusters:        clusterWatcher,
		SnapshotHistory: snapshotHistory,
		VerboseMode:     opts.VerboseMode,
		SettingsRef:     opts.SettingsRef,
	}

	eg, ctx := errgroup.WithContext(ctx)

	for name, start := range startFuncs {
		start := start // pike
		namedCtx := ctx
		if name != "" {
			namedCtx = contextutils.WithLogger(namedCtx, name)
		}
		eg.Go(func() error {
			contextutils.LoggerFrom(namedCtx).Debugf("starting main goroutine")
			if synced := params.MasterManager.GetCache().WaitForCacheSync(namedCtx); !synced {
				return eris.Errorf("caches failed to sync")
			}
			err := start(namedCtx, params)
			if err != nil {
				contextutils.LoggerFrom(namedCtx).Errorf("main goroutine failed: %v", err)
			}
			return err
		})
	}

	if clusterWatcher != nil {
		// start multicluster watches
		eg.Go(func() error {
			if synced := params.MasterManager.GetCache().WaitForCacheSync(ctx); !synced {
				return eris.Errorf("caches failed to sync")
			}
			return clusterWatcher.Run(mgr)
		})
	}

	eg.Go(func() error {
		// start the local manager
		ctx := contextutils.WithLogger(ctx, "controller-runtime-manager")
		contextutils.LoggerFrom(ctx).Infof("starting manager with options %+v", opts)
		return mgr.Start(ctx)
	})

	return eg.Wait()
}

// get the manager for the local cluster; we will use this as our "master" cluster
func makeMasterManager(opts Options, schemes runtime.SchemeBuilder) (manager.Manager, error) {
	cfg, err := config.GetConfigWithContext(opts.ManagementContext)
	if err != nil {
		return nil, err
	}

	mgr, err := manager.New(cfg, manager.Options{
		Namespace:          opts.MasterNamespace, // TODO (ilackarms): support configuring multiple watch namespaces on master cluster
		MetricsBindAddress: "0",                  // serve metrics using custom stats server
	})
	if err != nil {
		return nil, err
	}

	if schemes != nil {
		if err := schemes.AddToScheme(mgr.GetScheme()); err != nil {
			return nil, err
		}
	}
	return mgr, nil
}

func setupLogging(verboseMode, jsonLogging bool) {
	level := zapcore.InfoLevel
	if verboseMode {
		level = zapcore.DebugLevel
	}
	atomicLevel := zap.NewAtomicLevelAt(level)
	zapOpts := []zaputil.Opts{
		zaputil.Level(&atomicLevel),
	}
	if !jsonLogging {
		zapOpts = append(zapOpts,
			// Only set debug mode if specified. This will use a non-json (human readable) encoder which makes it impossible
			// to use any json parsing tools for the log. Should only be enabled explicitly
			zaputil.UseDevMode(true),
		)
	}
	baseLogger := zaputil.NewRaw(zapOpts...)

	// klog
	zap.ReplaceGlobals(baseLogger)

	// controller-runtime
	zapLogger := zapr.NewLogger(baseLogger)
	log.SetLogger(zapLogger)
	klog.SetLogger(zapLogger)

	// go-utils
	contextutils.SetFallbackLogger(baseLogger.Sugar())
}
