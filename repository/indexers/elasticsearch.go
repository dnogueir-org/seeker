package indexers

import "github.com/dnogueir-org/seeker/internal/entity"

type ElasticsearchIndexer struct{}

func (esi *ElasticsearchIndexer) Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) IndexResponse {

	return IndexResponse{
		Products: []string{"teste"},
		Page:     1,
		Total:    10,
	}
}
