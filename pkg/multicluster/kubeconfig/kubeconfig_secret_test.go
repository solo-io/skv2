package kubeconfig_test

import (
	"fmt"

	. "github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("KubeConfig Secret Conversions", func() {
	var (
		clusterName      = "test-name"
		kubeConfigFormat = `apiVersion: v1
clusters:
- cluster:
    server: test-server
  name: %s
contexts:
- context:
    cluster: test-name
    user: test-name
  name: test-name
current-context: test-name
kind: Config
preferences: {}
users:
- name: test-name
  user:
    token: alphanumericgarbage
`
		namespace     = "secret-namespace"
		kubeConfigRaw = fmt.Sprintf(kubeConfigFormat, clusterName)
		config        *api.Config
		err           error
	)

	BeforeEach(func() {
		config, err = clientcmd.Load([]byte(kubeConfigRaw))
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("ToSecret", func() {

		It("should convert a single KubeConfig to a single secret", func() {
			expectedLabels := map[string]string{
				"foo": "bar",
			}
			expectedSecret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      clusterName,
					Namespace: namespace,
					Labels:    expectedLabels,
				},
				Data: map[string][]byte{
					Key: []byte(kubeConfigRaw),
				},
				Type: SecretType,
			}
			secret, err := ToSecret(namespace, clusterName, expectedLabels, *config)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret).To(Equal(expectedSecret))
		})

	})

	Describe("SecretToConfig", func() {

		It("works", func() {
			secret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      clusterName,
					Namespace: namespace,
				},
				Data: map[string][]byte{
					Key: []byte(kubeConfigRaw),
				},
				Type: SecretType,
			}

			actualCluster, actualConfig, err := SecretToConfig(secret)
			Expect(err).NotTo(HaveOccurred())
			Expect(actualCluster).To(Equal(clusterName))
			Expect(actualConfig).NotTo(BeNil())
		})

	})

})
