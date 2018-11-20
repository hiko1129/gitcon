package infrastracture_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiko1129/gitcon/infrastracture"
)

var _ = Describe("ContributionClient", func() {
	var client *infrastracture.ContirubutionClient
	BeforeEach(func() {
		client, _ = infrastracture.NewContirubutionClient()
	})

	Describe("#FetchContributions", func() {
		Context("when valid username", func() {
			It("expects no error", func() {
				// confirmation real endpoint
				contributions, err := client.FetchContributions("hiko1129")
				Expect(contributions).NotTo(BeEmpty())
				Expect(err).NotTo(HaveOccurred())
			})
		})

		// Context("when invalid username", func() {
		// 	It("expects error", func() {
		// 		// confirmation real endpoint
		// 		contributions, err := client.FetchContributions("-01-01-01-")
		// 		Expect(contributions).To(BeEmpty())
		// 		Expect(err).To(HaveOccurred())
		// 	})
		// })
	})
})
