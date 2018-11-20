package usecase_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiko1129/gitcon/usecase"
)

type ContributionClient struct {
}

func (c *ContributionClient) FetchContributions(username string) (map[string]int, error) {
	if username == "hiko1129" {
		return map[string]int{"2018-11-09": 1, "2018-11-01": 4}, nil
	}

	return map[string]int{}, errors.New("Not Found")
}

var _ = Describe("FetchConstributions", func() {
	Describe("#Exec", func() {
		Context("when fetching contributions is successful", func() {
			It("returns contributions", func() {
				req := &usecase.FetchContributionsRequest{Username: "hiko1129"}
				u, _ := usecase.NewFetchContributions(req, &ContributionClient{})
				res, err := u.Exec()
				Expect(res.Contributions).To(Equal(map[string]int{"2018-11-09": 1, "2018-11-01": 4}))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when fetching contributions is failed", func() {
			It("expects error", func() {
				req := &usecase.FetchContributionsRequest{Username: "hogehoge"}
				u, _ := usecase.NewFetchContributions(req, &ContributionClient{})
				res, err := u.Exec()
				Expect(res.Contributions).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
