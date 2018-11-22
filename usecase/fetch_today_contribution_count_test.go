package usecase_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/hiko1129/gitcon/usecase"
)

type MutableContributionClient struct {
}

func (c *MutableContributionClient) FetchContributions(username string) (map[string]int, error) {
	if username == "hiko1129" {
		return map[string]int{time.Now().Format("2006-01-02"): 1}, nil
	}

	return map[string]int{}, errors.New("Not Found")
}

var _ = Describe("FetchTodayContributionCount", func() {
	Context("when fetching contributions is successful", func() {
		It("returns contributions", func() {
			req := &usecase.FetchTodayContributionCountRequest{Username: "hiko1129"}
			u := usecase.NewFetchTodayContributionCount(req, &MutableContributionClient{})
			res, err := u.Exec()
			Expect(res.Contribution).To(Equal(1))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when fetching contributions is failed", func() {
		It("returns contributions", func() {
			req := &usecase.FetchTodayContributionCountRequest{Username: "hogehoge"}
			u := usecase.NewFetchTodayContributionCount(req, &MutableContributionClient{})
			res, err := u.Exec()
			Expect(res.Contribution).To(Equal(0))
			Expect(err).To(HaveOccurred())
		})
	})
})
