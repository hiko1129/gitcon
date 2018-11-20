package usecase

import (
	"time"

	"github.com/hiko1129/gitcon/domain/client"
	"github.com/pkg/errors"
)

// FetchTodayContributionCountRequest struct
type FetchTodayContributionCountRequest struct {
	Username string
}

// FetchTodayContributionCountResponse struct
type FetchTodayContributionCountResponse struct {
	Contribution int
}

// FetchTodayContributionCount struct
type FetchTodayContributionCount struct {
	request *FetchTodayContributionCountRequest
	client  client.Contribution
}

// NewFetchTodayContributionCount func
func NewFetchTodayContributionCount(request *FetchTodayContributionCountRequest, client client.Contribution) (*FetchTodayContributionCount, error) {
	return &FetchTodayContributionCount{request: request, client: client}, nil
}

// Exec func
func (f *FetchTodayContributionCount) Exec() (*FetchTodayContributionCountResponse, error) {
	fe := &FetchTodayContributionCountResponse{}
	u, err := NewFetchContributionCount(&FetchContributionCountRequest{Username: f.request.Username}, f.client, time.Now())
	if err != nil {
		return fe, err
	}
	res, err := u.Exec()
	if err != nil {
		return fe, errors.Wrap(err, "fetch today contribution failed")
	}

	fe.Contribution = res.Contribution
	return fe, nil
}
