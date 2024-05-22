package source

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/solo-io/go-utils/contextutils"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Stoppable is a stoppable source
type Stoppable interface {
	source.Source
	InjectStopChannel(<-chan struct{}) error
}

// DynamicSource is a funnel for sources that can be
// dynamically (de)registered before & after the controller has started
type Dynamic interface {
	source.Source

	// sources must be registered with a unique id
	Add(id string, src source.Source) error

	// remove a source. errors if not found
	Remove(id string) error
}

// cache of sources
type cachedSource struct {
	// the original source
	source source.Source

	// cancel function to stop it
	cancel context.CancelFunc
}

// the args with which the dynamic source was started
type startArgs struct {
	i workqueue.RateLimitingInterface
}

// DynamicSource implements Dynamic
type DynamicSource struct {
	// cancel this context to stop all registered sources
	ctx context.Context

	// the cached sources that can be dynamically added/removed
	cache map[string]cachedSource

	// cache access
	lock sync.RWMutex

	// has source started?
	started *startArgs
}

func NewDynamicSource(ctx context.Context) *DynamicSource {
	return &DynamicSource{
		ctx:   ctx,
		cache: make(map[string]cachedSource),
	}
}

// start all the sources
func (s *DynamicSource) Start(ctx context.Context, i workqueue.RateLimitingInterface) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.started != nil {
		return errors.Errorf("source was already started")
	}

	for _, src := range s.cache {

		if err := src.source.Start(ctx, i); err != nil {
			return err
		}
	}

	s.started = &startArgs{
		i: i,
	}

	return nil
}

// only Stoppable sources are currently supported
func (s *DynamicSource) Add(id string, src Stoppable) error {
	contextutils.LoggerFrom(s.ctx).DPanic("DynamicSource.Add() may not work as expected due to the removal of dependency injection functions from controller-runtime in 15.0. See https://github.com/kubernetes-sigs/controller-runtime/releases")
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exists := s.cache[id]; exists {
		return errors.Errorf("source %v already exists", id)
	}

	ctx, cancel := context.WithCancel(s.ctx)
	if err := src.InjectStopChannel(ctx.Done()); err != nil {
		return err
	}

	if s.started != nil {
		if err := src.Start(ctx, s.started.i); err != nil {
			return errors.Wrapf(err, "failed to start source %v", id)
		}
	}

	s.cache[id] = cachedSource{
		source: src,
		cancel: cancel,
	}

	return nil
}

// remove (and stop) a source
func (s *DynamicSource) Remove(id string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	src, ok := s.cache[id]
	if !ok {
		return errors.Errorf("no source in cache with id %v", id)
	}

	src.cancel()

	delete(s.cache, id)

	return nil
}
