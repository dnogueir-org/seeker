package indexers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dnogueir-org/seeker/internal/entity"
	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchIndexer struct {
	ElasticClient *elasticsearch.Client
	IndexName     string
}

func (esi *ElasticsearchIndexer) Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) IndexResponseHits {

	response := esi.getDocumentsByFields(page, resultsPerPage, fields)

	return response
}

func (esi *ElasticsearchIndexer) getDocumentsByFields(page, resultsPerPage int, fields []entity.Field) IndexResponseHits {

	query := `{
		"query":{
			"bool":{
				 "must":[
						{
							 "term":{
									"filters.genero.value":{
										 "value":"masculino"
									}
							 }
						}
				 ]
			}
	 }
	}`

	res, err := esi.ElasticClient.Search(
		esi.ElasticClient.Search.WithContext(context.Background()),
		esi.ElasticClient.Search.WithIndex(esi.IndexName),
		esi.ElasticClient.Search.WithBody(strings.NewReader(query)),
		esi.ElasticClient.Search.WithSize(resultsPerPage),
		esi.ElasticClient.Search.WithTrackTotalHits(true),
		esi.ElasticClient.Search.WithPretty(),
	)
	if err != nil {
		fmt.Println(err.Error())
		return IndexResponseHits{}
	}
	defer res.Body.Close()

	var (
		response IndexResponseHits
		hits     IndexHits
	)

	hits.Hits = &response

	if err := json.NewDecoder(res.Body).Decode(&hits); err != nil {
		fmt.Println(err.Error())
		return IndexResponseHits{}
	}

	return response
}
