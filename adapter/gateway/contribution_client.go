package gateway

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type contributionsPerUser map[string]map[string]int

// ContirubutionClient struct
type ContirubutionClient struct {
	http          HTTP
	contributions contributionsPerUser
}

// NewContirubutionClient func
func NewContirubutionClient(http HTTP) *ContirubutionClient {
	return &ContirubutionClient{http: http}
}

const (
	userEndpoint = "https://github.com/users/"
)

// FetchContributions func
func (c *ContirubutionClient) FetchContributions(username string) (map[string]int, error) {
	if c.hasMemo(username) {
		return c.getContributionsCache(username), nil
	}

	result := map[string]int{}
	res, err := c.http.Get(userEndpoint + username + "/contributions")
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return result, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	contributions, err := c.extractContributions(res.Body)

	c.setMemo(username, contributions)
	return contributions, nil
}

func (c *ContirubutionClient) extractContributions(body io.ReadCloser) (map[string]int, error) {
	contributions := map[string]int{}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return contributions, err
	}

	doc.Find(".js-calendar-graph-svg").Find("g").Children().Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(i int, s *goquery.Selection) {
			count, ok := s.Attr("data-count")
			if !ok {
				err = errors.New("data-count attribute does not exsist")
			}
			date, ok := s.Attr("data-date")
			if !ok {
				err = errors.New("data-date attribute does not exsist")
			}
			c, _ := strconv.Atoi(count)
			contributions[date] = c
		})
	})

	if err != nil {
		return contributions, err
	}

	return contributions, nil
}

func (c *ContirubutionClient) hasMemo(username string) bool {
	_, ok := c.contributions[username]
	return ok
}

func (c *ContirubutionClient) setMemo(username string, contributions map[string]int) {
	if c.contributions == nil {
		c.contributions = contributionsPerUser{username: contributions}
		return
	}
	c.contributions[username] = contributions
}

func (c *ContirubutionClient) getContributionsCache(username string) map[string]int {
	return c.contributions[username]
}
