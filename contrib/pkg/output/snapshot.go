package output

import (
	"context"
	"fmt"

	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	"github.com/solo-io/skv2/pkg/multicluster"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// the key used to differentiate discovery resources by
// the cluster in which they were discovered
var ClusterLabelKey = fmt.Sprintf("cluster.%s", v1alpha1.SchemeGroupVersion.Group)

// adds cluster labels to the given set of labels and returns them
func AddClusterLabels(clusterName string, objLabels map[string]string) map[string]string {
	// add cluster label to object
	if objLabels == nil {
		objLabels = map[string]string{}
	}
	for k, v := range ClusterLabels(clusterName) {
		objLabels[k] = v
	}
	return objLabels
}

// Create a label that identifies the cluster used to discover a resource.
func ClusterLabels(cluster string) map[string]string {
	clusterK, clusterV := ClusterLabel(cluster)
	return map[string]string{
		clusterK: clusterV,
	}
}

func ClusterLabel(cluster string) (string, string) {
	return ClusterLabelKey,
		fmt.Sprintf("%s", cluster)
}

// User-defined error-handling functions.
// Used to invoke custom error-handling code when a resource write fails.
type ErrorHandler interface {
	// handle an error that happens when we try to write a resource
	HandleWriteError(resource ezkube.Object, err error)

	// handle an error that happens when we try to delete a resource
	HandleDeleteError(resource ezkube.Object, err error)

	// handle an error that happens when we try list resources in a cluster
	HandleListError(err error)
}

// straightforward implementation of an Error Handler
type ErrorHandlerFuncs struct {
	// handle an error that happens when we try to write a resource
	HandleWriteErrorFunc func(resource ezkube.Object, err error)

	// handle an error that happens when we try to delete a resource
	HandleDeleteErrorFunc func(resource ezkube.Object, err error)

	// handle an error that happens when we try list resources in a cluster
	HandleListErrorFunc func(err error)
}

func (e ErrorHandlerFuncs) HandleWriteError(resource ezkube.Object, err error) {
	if e.HandleWriteErrorFunc != nil {
		e.HandleWriteErrorFunc(resource, err)
	}
}

func (e ErrorHandlerFuncs) HandleDeleteError(resource ezkube.Object, err error) {
	if e.HandleDeleteErrorFunc != nil {
		e.HandleDeleteErrorFunc(resource, err)
	}
}

func (e ErrorHandlerFuncs) HandleListError(err error) {
	if e.HandleListErrorFunc != nil {
		e.HandleListErrorFunc(err)
	}
}

var (
	// resources_synced_total holds the total number of times a resource is synced successfully to storage.
	resourcesSyncedTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "resources_synced_total",
		Help: "Total number of successful resource writes to storage. result indicates the result of the write, i.e. created, updated, unchanged",
	}, []string{"snapshot", "result", "type", "ref"})

	incrementResourcesSyncedTotal = func(snapshot, result string, obj ezkube.Object) {
		resourcesSyncedTotal.WithLabelValues(
			snapshot,
			fmt.Sprintf("%T", obj),
			obj.GetNamespace()+"/"+obj.GetName(),
			string(result),
		).Inc()
	}

	// resources_deleted_total holds the total number of times a resource is deleted from storage.
	resourcesDeletedTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "resources_deleted_total",
		Help: "Total number of successful resource deletes to storage.",
	}, []string{"snapshot", "type", "ref"})

	incrementResourcesDeletedTotal = func(snapshot string, obj ezkube.Object) {
		resourcesDeletedTotal.WithLabelValues(
			snapshot,
			fmt.Sprintf("%T", obj),
			obj.GetNamespace()+"/"+obj.GetName(),
		).Inc()
	}

	// resource_write_fails_total holds the total number of failed attempts to write to storage (either as a delete or create)
	resourcesWriteFailsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "resource_write_fails_total",
		Help: "Total number of failures encountered when attempting to write a resource to storage. action indicates whether this was an upsert or a delete",
	}, []string{"snapshot", "type", "ref", "action"})

	incrementResourcesWriteFailsTotal = func(snapshot, action string, obj ezkube.Object) {
		resourcesWriteFailsTotal.WithLabelValues(
			snapshot,
			fmt.Sprintf("%T", obj),
			obj.GetNamespace()+"/"+obj.GetName(),
			action,
		).Inc()
	}
)

// This util package helps with syncing batches of resources in one place.
// It uses labels to reconcile the diff between the existing state
// and handles manual garbage collection of resources which
// for one reason or another cannot be garbage collected.

func init() {
	metrics.Registry.MustRegister(
		resourcesSyncedTotal,
		resourcesDeletedTotal,
		resourcesWriteFailsTotal,
	)
}

// a ResourceList define a list of resources we wish to write to
// kubernetes. A ListFunc can be provided to compare the resources
// with what is currently written to storage and trim stale resources.
// A transition function can also be provided for updating existing resources.
type ResourceList struct {
	// the desired resources should share the given labels
	Resources []ezkube.Object

	// list function that will be used to compare the given resources to the currently existing state.
	// if the resources in the list do not match those returned by the list func,
	// the differences will be reconciled by applying the snapshot.
	// if this function is nil, no garbage collection will be done.
	ListFunc func(ctx context.Context, cli client.Client) ([]ezkube.Object, error)

	// name of resource Kind, used for debugging only
	ResourceKind string
}

// partition the resource list by the ClusterName of each object.
func (l ResourceList) SplitByClusterName() map[string][]ezkube.Object {
	listsByCluster := map[string][]ezkube.Object{}
	for _, resource := range l.Resources {
		clusterName := resource.GetClusterName()
		listForCluster := listsByCluster[clusterName]

		// list func (i.e. garbage collection labels) shared across clusters

		listForCluster = append(listForCluster, resource)

		listsByCluster[clusterName] = listForCluster
	}
	return listsByCluster
}

// an Output Snapshot defines a list of desired resources
// to apply to Kubernetes.
// Stale resources (resources with no parents) will be garbage collected.
//
// A resources is determined to be stale when it currently exists
// in the cluster, but does not exist in the snapshot.
type Snapshot struct {
	// name of the snapshot, used for metrics
	Name string

	ListsToSync []ResourceList
}

// sync the output snapshot to local cluster storage.
// only writes resources intended for the local cluster (with ClusterName == "")
// Note that Syncing snapshots in this way adds the label
func (s Snapshot) SyncLocalCluster(ctx context.Context, cli client.Client, errHandler ErrorHandler) {

	for _, list := range s.ListsToSync {
		listForLocalCluster := list.SplitByClusterName()[multicluster.LocalCluster]

		resourcesForLocalCluster := ResourceList{
			Resources:    listForLocalCluster,
			ListFunc:     list.ListFunc,
			ResourceKind: list.ResourceKind,
		}

		s.syncList(ctx, multicluster.LocalCluster, cli, resourcesForLocalCluster, errHandler)
	}

}

// sync the output snapshot to storage across multiple clusters.
// uses the object's ClusterName to determine the correct destination cluster.
func (s Snapshot) SyncMultiCluster(ctx context.Context, mcClient multicluster.Client, errHandler ErrorHandler) {

	clusters := mcClient.ListClusters()
	for _, list := range s.ListsToSync {
		listsByCluster := list.SplitByClusterName()
		// TODO(ilackarms): possible error case that we're ignoring here;
		// we only write resources to clusters that are available to the multicluster client
		// if the cluster is not available, we will not error (simply skip writing the resources here)
		for _, cluster := range clusters {
			listForCluster := listsByCluster[cluster]

			cli, err := mcClient.Cluster(cluster)
			if err != nil {
				for _, object := range listForCluster {
					// send a write error for every object in the cluster
					errHandler.HandleWriteError(object, err)
				}
				continue
			}

			resourcesForCluster := ResourceList{
				Resources:    listForCluster,
				ListFunc:     list.ListFunc,
				ResourceKind: list.ResourceKind,
			}

			s.syncList(ctx, cluster, cli, resourcesForCluster, errHandler)
		}
	}
}

// sync the list to the cluster.
// clientCluster represents the cluster to which the resources are being written.
// garbage collects every resource turned by the list.ListFunc which is no longer desired.
// any resources written for a different clientCluster will not be garbage collected.
func (s Snapshot) syncList(ctx context.Context, clientCluster string, cli client.Client, list ResourceList, errHandler ErrorHandler) {
	for _, obj := range list.Resources {

		// upsert all desired resources
		if err := s.upsert(ctx, cli, obj); err != nil {
			errHandler.HandleWriteError(obj, err)
		}
	}

	if list.ListFunc == nil {
		return
	}

	// remove stale resources
	existingList, err := list.ListFunc(ctx, cli)
	if err != nil {
		// cache read should never error
		contextutils.LoggerFrom(ctx).Errorf("failed to read from cache for kind %v: %v", list.ResourceKind, err)
		return
	}

	isStale := func(res metav1.Object) bool {
		// only process resources with the cluster label
		// matching the resources written for this cluster.
		// this is used to distinguish between resources written to
		// a local cluster which is also registered as a managed cluster.
		for k, v := range ClusterLabels(clientCluster) {
			if res.GetLabels()[k] != v {
				return false
			}
		}
		for _, desired := range list.Resources {
			if res.GetName() == desired.GetName() && res.GetNamespace() == desired.GetNamespace() {
				return false
			}
		}
		return true
	}

	var staleRess []ezkube.Object
	for _, existingRes := range existingList {
		existingRes := existingRes // pike
		if isStale(existingRes) {
			staleRess = append(staleRess, existingRes)
		}
	}

	for _, obj := range staleRess {
		if err := s.delete(ctx, cli, obj); err != nil {
			errHandler.HandleDeleteError(obj, err)
		}
	}
}

func (s Snapshot) upsert(ctx context.Context, cli client.Client, obj ezkube.Object) error {
	// add cluster label to object
	obj.SetLabels(
		AddClusterLabels(obj.GetClusterName(), obj.GetLabels()),
	)

	result, err := controllerutils.Upsert(ctx, cli, obj)
	if err != nil {
		contextutils.LoggerFrom(ctx).Errorw("failed upserting resource", "resource", sets.TypedKey(obj), "err", err)

		incrementResourcesWriteFailsTotal(s.Name, "upsert", obj)

		return err
	}
	contextutils.LoggerFrom(ctx).Debugw("upserted resource", "resource", sets.TypedKey(obj))

	incrementResourcesSyncedTotal(
		s.Name,
		string(result),
		obj,
	)

	return nil
}

func (s Snapshot) delete(ctx context.Context, cli client.Client, obj ezkube.Object) error {
	if err := cli.Delete(ctx, obj); err != nil {
		contextutils.LoggerFrom(ctx).Errorw("failed deleting stale resource", "resource", sets.TypedKey(obj), "err", err)

		incrementResourcesWriteFailsTotal(s.Name, "delete", obj)

		return err
	}
	contextutils.LoggerFrom(ctx).Debugw("deleted resource", "resource", sets.TypedKey(obj))

	incrementResourcesDeletedTotal(s.Name, obj)
	return nil
}
