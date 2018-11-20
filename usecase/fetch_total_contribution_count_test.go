package usecase_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiko1129/gitcon/usecase"
)

var _ = Describe("FetchTotalContributionCount", func() {
	Context("when fetching contributions is successful", func() {
		It("returns contribution", func() {
			req := &usecase.FetchTotalContributionCountRequest{Username: "hiko1129"}
			u, _ := usecase.NewFetchTotalContributionCount(req, &ContributionClient{})
			res, err := u.Exec()
			Expect(res.Contribution).To(Equal(5))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when fetching contributions is failed", func() {
		It("expects error", func() {
			req := &usecase.FetchTotalContributionCountRequest{Username: "hoge"}
			u, _ := usecase.NewFetchTotalContributionCount(req, &ContributionClient{})
			res, err := u.Exec()
			Expect(res.Contribution).To(Equal(0))
			Expect(err).To(HaveOccurred())
		})
	})
})
