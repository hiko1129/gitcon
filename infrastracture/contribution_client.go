package infrastracture

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// ContirubutionClient struct
type ContirubutionClient struct {
	username      string
	contributions map[string]int
}

// NewContirubutionClient func
func NewContirubutionClient() (*ContirubutionClient, error) {
	return &ContirubutionClient{}, nil
}

const (
	userEndpoint = "https://github.com/users/"
)

// FetchContributions func
func (c *ContirubutionClient) FetchContributions(username string) (map[string]int, error) {
	if c.hasMemo(username) {
		return c.getContributionsCache(), nil
	}

	result := map[string]int{}
	res, err := http.Get(userEndpoint + username + "/contributions")
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return result, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return result, err
	}

	contributions := map[string]int{}
	doc.Find(".js-calendar-graph-svg").Find("g").Children().Each(func(i int, s *goquery.Selection) {
		if s.Find("g") != nil {
			s.Children().Each(func(i int, s *goquery.Selection) {
				count, ok := s.Attr("data-count")
				if !ok {
					log.Fatal("data-count attribute does not exsist")
				}
				date, ok := s.Attr("data-date")
				if !ok {
					log.Fatal("data-date attribute does not exsist")
				}
				c, _ := strconv.Atoi(count)
				contributions[date] = c
			})
		} else {
			log.Fatal("maybe web page changed")
		}
	})

	c.setMemo(username, contributions)
	return contributions, nil
}

func (c *ContirubutionClient) hasMemo(username string) bool {
	return c.username == username && len(c.contributions) != 0
}

func (c *ContirubutionClient) setMemo(username string, contributions map[string]int) {
	c.username = username
	c.contributions = contributions
}

func (c *ContirubutionClient) getContributionsCache() map[string]int {
	return c.contributions
}
