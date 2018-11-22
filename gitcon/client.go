package gitcon

import (
	"time"

	"github.com/hiko1129/gitcon/adapter/gateway"
	"github.com/hiko1129/gitcon/infrastructure"
	"github.com/hiko1129/gitcon/usecase"
)

// Client struct
type Client struct {
	username string
	client   *gateway.ContirubutionClient
}

// New func
func New(username string) (*Client, error) {
	http := infrastructure.NewHTTP()
	client := gateway.NewContirubutionClient(http)

	return &Client{username: username, client: client}, nil
}

// FetchTodayContributionCount func
func (c *Client) FetchTodayContributionCount() (int, error) {
	req := &usecase.FetchTodayContributionCountRequest{Username: c.username}
	u := usecase.NewFetchTodayContributionCount(req, c.client)
	res, err := u.Exec()
	if err != nil {
		return 0, err
	}

	return res.Contribution, nil
}

// FetchContributionCount func
func (c *Client) FetchContributionCount(datetime time.Time) (int, error) {
	req := &usecase.FetchContributionCountRequest{Username: c.username}
	u := usecase.NewFetchContributionCount(req, c.client, datetime)
	res, err := u.Exec()
	if err != nil {
		return 0, err
	}

	return res.Contribution, nil
}

// FetchContributions func
func (c *Client) FetchContributions() (map[string]int, error) {
	result := map[string]int{}
	req := &usecase.FetchContributionsRequest{Username: c.username}
	u := usecase.NewFetchContributions(req, c.client)
	res, err := u.Exec()
	if err != nil {
		return result, err
	}

	return res.Contributions, nil
}

// FetchTotalContributionCount func
func (c *Client) FetchTotalContributionCount() (int, error) {
	req := &usecase.FetchTotalContributionCountRequest{Username: c.username}
	u := usecase.NewFetchTotalContributionCount(req, c.client)
	res, err := u.Exec()
	if err != nil {
		return 0, err
	}

	return res.Contribution, nil
}
