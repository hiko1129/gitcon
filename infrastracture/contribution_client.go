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

	return contributions, nil
}
