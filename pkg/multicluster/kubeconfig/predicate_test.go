package kubeconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

var _ = Describe("KubeConfig Secret Predicate", func() {
	Describe("BuildPredicate", func() {

		It("does not process events for non-secrets", func() {
			watchNamespaces := []string{}
			pred := BuildPredicate(watchNamespaces)

			obj := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "configmap1",
					Namespace: "ns1",
				},
				Data: map[string]string{
					"a": "b",
				},
			}
			Expect(pred.CreateFunc(event.CreateEvent{Object: obj})).To(BeFalse())
			Expect(pred.DeleteFunc(event.DeleteEvent{Object: obj})).To(BeFalse())
			Expect(pred.UpdateFunc(event.UpdateEvent{ObjectNew: obj})).To(BeFalse())
			Expect(pred.GenericFunc(event.GenericEvent{Object: obj})).To(BeFalse())
		})

		It("does not process events for non-kubeconfig secrets", func() {
			watchNamespaces := []string{}
			pred := BuildPredicate(watchNamespaces)

			obj := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns1",
				},
				Type: corev1.SecretTypeOpaque,
			}
			Expect(pred.CreateFunc(event.CreateEvent{Object: obj})).To(BeFalse())
			Expect(pred.DeleteFunc(event.DeleteEvent{Object: obj})).To(BeFalse())
			Expect(pred.UpdateFunc(event.UpdateEvent{ObjectNew: obj})).To(BeFalse())
			Expect(pred.GenericFunc(event.GenericEvent{Object: obj})).To(BeFalse())
		})

		It("processes events in all namespaces when watchNamespaces is empty", func() {
			watchNamespaces := []string{}
			pred := BuildPredicate(watchNamespaces)

			objOld := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns1",
				},
				Type: SecretType,
				Data: map[string][]byte{
					"a": []byte("b"),
				},
			}
			objNew := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns1",
				},
				Type: SecretType,
				Data: map[string][]byte{
					"c": []byte("d"),
				},
			}
			Expect(pred.CreateFunc(event.CreateEvent{Object: objNew})).To(BeTrue())
			Expect(pred.DeleteFunc(event.DeleteEvent{Object: objNew})).To(BeTrue())
			Expect(pred.UpdateFunc(event.UpdateEvent{ObjectOld: objOld, ObjectNew: objNew})).To(BeTrue())
			Expect(pred.GenericFunc(event.GenericEvent{Object: objNew})).To(BeTrue())
		})

		It("processes events only in watchNamespaces", func() {
			watchNamespaces := []string{"ns1", "ns2"}
			pred := BuildPredicate(watchNamespaces)

			objOld := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns2",
				},
				Type: SecretType,
				Data: map[string][]byte{
					"a": []byte("b"),
				},
			}
			objNew := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns2",
				},
				Type: SecretType,
				Data: map[string][]byte{
					"c": []byte("d"),
				},
			}
			Expect(pred.CreateFunc(event.CreateEvent{Object: objNew})).To(BeTrue())
			Expect(pred.DeleteFunc(event.DeleteEvent{Object: objNew})).To(BeTrue())
			Expect(pred.UpdateFunc(event.UpdateEvent{ObjectOld: objOld, ObjectNew: objNew})).To(BeTrue())
			Expect(pred.GenericFunc(event.GenericEvent{Object: objNew})).To(BeTrue())

			// change to a non-watched namespace
			objOld.ObjectMeta.Namespace = "ns3"
			objNew.ObjectMeta.Namespace = "ns3"
			Expect(pred.CreateFunc(event.CreateEvent{Object: objNew})).To(BeFalse())
			Expect(pred.DeleteFunc(event.DeleteEvent{Object: objNew})).To(BeFalse())
			Expect(pred.UpdateFunc(event.UpdateEvent{ObjectOld: objOld, ObjectNew: objNew})).To(BeFalse())
			Expect(pred.GenericFunc(event.GenericEvent{Object: objNew})).To(BeFalse())
		})

		It("does not process update events when secret didn't change", func() {
			watchNamespaces := []string{}
			pred := BuildPredicate(watchNamespaces)

			objOld := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns1",
				},
				Type: SecretType,
				Data: map[string][]byte{
					"a": []byte("b"),
				},
			}
			objNew := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret1",
					Namespace: "ns1",
				},
				Type: SecretType,
				Data: map[string][]byte{
					"a": []byte("b"),
				},
			}
			Expect(pred.UpdateFunc(event.UpdateEvent{ObjectOld: objOld, ObjectNew: objNew})).To(BeFalse())
		})
	})
})
