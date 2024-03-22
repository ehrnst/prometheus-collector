package crds

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

// var K8sClient 	*kubernetes.Clientset
// var Cfg       	*rest.Config

func TestCRDs(t *testing.T) {
  RegisterFailHandler(Fail)

  RunSpecs(t, "CRDs Test Suite")
}

var _ = BeforeSuite(func() {
  // var err error
  // K8sClient, Cfg, err = utils.SetupKubernetesClient()
  // Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
  By("tearing down the test environment")
})
