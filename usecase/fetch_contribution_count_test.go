package usecase_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiko1129/gitcon/usecase"
)

var _ = Describe("FetchContributionCount", func() {
	Describe("#Exec", func() {
		Context("when fetching contributions is successful", func() {
			Context("when there is corresponding date", func() {
				It("returns contributions", func() {
					req := &usecase.FetchContributionCountRequest{Username: "hiko1129"}
					loc, _ := time.LoadLocation("Asia/Tokyo")
					u, _ := usecase.NewFetchContributionCount(req, &ContributionClient{}, time.Date(2018, 11, 9, 0, 0, 0, 0, loc))
					res, err := u.Exec()
					Expect(res.Contribution).To(Equal(1))
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when there isn't corresponding date", func() {
				It("returns contributions", func() {
					req := &usecase.FetchContributionCountRequest{Username: "hiko1129"}
					loc, _ := time.LoadLocation("Asia/Tokyo")
					u, _ := usecase.NewFetchContributionCount(req, &ContributionClient{}, time.Date(2018, 11, 8, 0, 0, 0, 0, loc))
					res, err := u.Exec()
					Expect(res.Contribution).To(Equal(0))
					Expect(err).To(MatchError("no corresponding date"))
				})
			})
		})

		Context("when fetching contributions is failed", func() {
			It("expects error", func() {
				req := &usecase.FetchContributionCountRequest{Username: "hoge"}
				u, _ := usecase.NewFetchContributionCount(req, &ContributionClient{}, time.Now())
				res, err := u.Exec()
				Expect(res.Contribution).To(Equal(0))
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
