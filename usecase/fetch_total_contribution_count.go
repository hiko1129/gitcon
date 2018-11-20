package usecase

import (
	"github.com/hiko1129/gitcon/domain/client"
	"github.com/pkg/errors"
)

// FetchTotalContributionCountRequest struct
type FetchTotalContributionCountRequest struct {
	Username string
}

// FetchTotalContributionCountResponse struct
type FetchTotalContributionCountResponse struct {
	Contribution int
}

// FetchTotalContributionCount struct
type FetchTotalContributionCount struct {
	request *FetchTotalContributionCountRequest
	client  client.Contribution
}

// NewFetchTotalContributionCount func
func NewFetchTotalContributionCount(request *FetchTotalContributionCountRequest, client client.Contribution) (*FetchTotalContributionCount, error) {
	return &FetchTotalContributionCount{request: request, client: client}, nil
}

// Exec func
func (f *FetchTotalContributionCount) Exec() (*FetchTotalContributionCountResponse, error) {
	fe := &FetchTotalContributionCountResponse{}
	contributions, err := f.client.FetchContributions(f.request.Username)
	if err != nil {
		return fe, errors.Wrap(err, "fetch total contribution failed")
	}

	sum := 0
	for _, contribution := range contributions {
		sum += contribution
	}
	fe.Contribution = sum

	return fe, nil
}
