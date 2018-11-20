package client

// Contribution interface
type Contribution interface {
	FetchContributions(string) (map[string]int, error)
}
