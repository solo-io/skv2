package access_test

import (
	"fmt"

	. "github.com/solo-io/skv2/pkg/multicluster/access"
	"k8s.io/client-go/tools/clientcmd/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("KubeConfig Secret Conversions", func() {
	var (
		clusterName      = "test-name1"
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
		kubeConfigRaw = fmt.Sprintf(kubeConfigFormat, clusterName)
		config        *api.Config
		err           error
	)

	BeforeEach(func() {
		config, err = clientcmd.Load([]byte(kubeConfigRaw))
		Expect(err).NotTo(HaveOccurred())
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
					clusterName: []byte(kubeConfigRaw),
				},
				Type: KubeConfigSecretType,
			}
			secret, err := KubeConfigToSecret(name, namespace, clusterName, config)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret).To(Equal(expectedSecret))
		})
	})
})
