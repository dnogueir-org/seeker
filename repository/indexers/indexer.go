package indexers

import "github.com/dnogueir-org/seeker/internal/entity"

type Indexer interface {
	Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) IndexResponse
}

type IndexResponse struct {
	Products []string
	Page     int
	Total    int
}
