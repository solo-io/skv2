package errhandlers

import (
	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
)

// this file contains ErrorHandlers for handling errors created when writing output snapshots

func ResourceWriteError(resource ezkube.Object, err error) error {
	return eris.Wrapf(err, "writing resource %v failed", sets.Key(resource))
}
func ResourceDeleteError(resource ezkube.Object, err error) error {
	return eris.Wrapf(err, "deleting resource %v failed", sets.Key(resource))
}
func ListError(err error) error {
	return eris.Wrapf(err, "listing failed")
}

type AppendingErrHandler struct {
	errs error
}

func (a *AppendingErrHandler) HandleWriteError(resource ezkube.Object, err error) {
	a.errs = multierror.Append(a.errs, ResourceWriteError(resource, err))
}

func (a *AppendingErrHandler) HandleDeleteError(resource ezkube.Object, err error) {
	a.errs = multierror.Append(a.errs, ResourceDeleteError(resource, err))
}

func (a *AppendingErrHandler) HandleListError(err error) {
	a.errs = multierror.Append(a.errs, ListError(err))
}

// Errors returns the errors collected by the handler
func (a *AppendingErrHandler) Errors() error {
	return a.errs
}
