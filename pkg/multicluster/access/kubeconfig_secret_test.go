package access_test

import (
	"fmt"

	. "github.com/solo-io/skv2/pkg/multicluster/access"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("KubeConfig Secret Conversions", func() {
	var (
		clusterName1  = "test-name1"
		kubeConfigRaw = `apiVersion: v1
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
		kubeConfigRaw1 = fmt.Sprintf(kubeConfigRaw, clusterName1)
		kubeConfig1    KubeConfig
	)

	BeforeEach(func() {
		config1, err := clientcmd.Load([]byte(kubeConfigRaw1))
		Expect(err).NotTo(HaveOccurred())
		kubeConfig1 = KubeConfig{
			Config:  *config1,
			Cluster: clusterName1,
		}
	})

	Describe("KubeConfigToSecret", func() {
		It("should convert a single KubeConfig to a single secret", func() {
			name := "secret-name"
			namespace := "secret-namespace"
			expectedSecret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Data: map[string][]byte{
					clusterName1: []byte(kubeConfigRaw1),
				},
				Type: KubeConfigSecretType,
			}
			secret, err := KubeConfigToSecret(name, namespace, &kubeConfig1)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret).To(Equal(expectedSecret))
		})
	})
})
