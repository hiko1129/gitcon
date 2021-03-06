package usecase

import (
	"github.com/hiko1129/gitcon/domain/client"
	"github.com/pkg/errors"
)

// FetchContributionsRequest struct
type FetchContributionsRequest struct {
	Username string
}

// FetchContributionsResponse struct
type FetchContributionsResponse struct {
	Contributions map[string]int
}

// FetchContributions struct
type FetchContributions struct {
	request *FetchContributionsRequest
	client  client.Contribution
}

// NewFetchContributions func
func NewFetchContributions(request *FetchContributionsRequest, client client.Contribution) *FetchContributions {
	return &FetchContributions{request: request, client: client}
}

// Exec func
func (f *FetchContributions) Exec() (*FetchContributionsResponse, error) {
	fe := &FetchContributionsResponse{}
	contributions, err := f.client.FetchContributions(f.request.Username)
	if err != nil {
		return fe, errors.Wrap(err, "fetch contributions failed")
	}

	fe.Contributions = contributions

	return fe, nil
}
