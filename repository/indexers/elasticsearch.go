package indexers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dnogueir-org/seeker/internal/entity"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchIndexer struct {
	ElasticClient *elasticsearch.Client
	IndexName     string
}

func (esi *ElasticsearchIndexer) Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) IndexResponse {

	req := esapi.GetRequest{
		Index:      esi.IndexName,
		DocumentID: "01399016",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := req.Do(ctx, esi.ElasticClient)
	if err != nil {
		fmt.Println(err.Error())
		return IndexResponse{}
	}
	defer res.Body.Close()

	var (
		response IndexResponse
		body     document
	)
	body.Source = &response

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return IndexResponse{}
	}

	return response
}
