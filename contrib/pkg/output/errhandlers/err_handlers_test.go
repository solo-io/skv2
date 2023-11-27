package errhandlers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/contrib/pkg/output"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Error Handlers", func() {
	var (
		errorHandler output.ErrorHandler
		fakeResource = &v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "fake-name",
				Namespace: "fake-ns",
			},
		}
	)

	Context("AppendingErrHandler", func() {
		BeforeEach(func() {
			errorHandler = &AppendingErrHandler{}
		})

		DescribeTable("should append errors to the error handler object", func(mutation errorMutation, expects errorExpectation) {
			mutation(errorHandler, fakeResource)
			expects(errorHandler.(*AppendingErrHandler).Errors(), fakeResource)
		},
			Entry("when HandleListError is called", addListError, expectListError),
			Entry("when HandleWriteError is called", addWriteError, expectWriteError),
			Entry("when HandleDeleteError is called", addDeleteError, expectDeleteError),
			Entry("when multiple error handler functions are invoked", func(handler output.ErrorHandler, resource ezkube.Object) {
				addListError(handler, resource)
				addWriteError(handler, resource)
				addDeleteError(handler, resource)
			}, func(err error, resource ezkube.Object) {
				Expect(err).To(And(
					MatchError(ListError(err)),
					MatchError(ResourceWriteError(resource, err)),
					MatchError(ResourceDeleteError(resource, err)),
				))
			}),
		)

		It("should not have an error when no errors occur", func() {
			errs := errorHandler.(*AppendingErrHandler).Errors()
			Expect(errs).NotTo(HaveOccurred())
		})
	})
})

type errorMutation func(handler output.ErrorHandler, resource ezkube.Object)

func addListError(handler output.ErrorHandler, _ ezkube.Object) {
	handler.HandleListError(errors.New(""))
}
func addWriteError(handler output.ErrorHandler, resource ezkube.Object) {
	handler.HandleWriteError(resource, errors.New(""))
}
func addDeleteError(handler output.ErrorHandler, resource ezkube.Object) {
	handler.HandleDeleteError(resource, errors.New(""))
}

// errorExpectation is a function that expects a specific error to occur
type errorExpectation func(err error, resource ezkube.Object)

func expectListError(err error, resource ezkube.Object) {
	Expect(err).To(HaveOccurred())
	Expect(err).To(And(
		MatchError(ListError(err)),
		Not(MatchError(ResourceWriteError(resource, err))),
		Not(MatchError(ResourceDeleteError(resource, err))),
	))
}
func expectWriteError(err error, resource ezkube.Object) {
	Expect(err).To(HaveOccurred())
	Expect(err).To(And(
		MatchError(ResourceWriteError(resource, err)),
		Not(MatchError(ListError(err))),
		Not(MatchError(ResourceDeleteError(resource, err))),
	))
}
func expectDeleteError(err error, resource ezkube.Object) {
	Expect(err).To(HaveOccurred())
	Expect(err).To(And(
		MatchError(ResourceDeleteError(resource, err)),
		Not(MatchError(ListError(err))),
		Not(MatchError(ResourceWriteError(resource, err))),
	))
}
