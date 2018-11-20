package usecase

import (
	"time"

	"github.com/hiko1129/gitcon/domain/client"
	"github.com/pkg/errors"
)

// FetchContributionCountRequest struct
type FetchContributionCountRequest struct {
	Username string
}

// FetchContributionCountResponse struct
type FetchContributionCountResponse struct {
	Contribution int
}

// FetchContributionCount struct
type FetchContributionCount struct {
	request  *FetchContributionCountRequest
	client   client.Contribution
	datetime time.Time
}

// NewFetchContributionCount func
func NewFetchContributionCount(request *FetchContributionCountRequest, client client.Contribution, datetime time.Time) (*FetchContributionCount, error) {
	return &FetchContributionCount{request: request, client: client, datetime: datetime}, nil
}

// Exec func
func (f *FetchContributionCount) Exec() (*FetchContributionCountResponse, error) {
	fe := &FetchContributionCountResponse{}
	contributions, err := f.client.FetchContributions(f.request.Username)
	if err != nil {
		return fe, errors.Wrap(err, "fetch contribution failed")
	}

	if val, ok := contributions[f.datetime.Format("2006-01-02")]; ok {
		fe.Contribution = val
		return fe, nil
	}

	return fe, errors.New("no corresponding date")
}
