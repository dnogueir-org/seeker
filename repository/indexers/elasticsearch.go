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

func (esi *ElasticsearchIndexer) Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) []IndexResponse {

	//response := esi.getDocById("01399016")
	response := esi.getDocumentsByFields(page, resultsPerPage, fields)

	return response
}

func (esi *ElasticsearchIndexer) getDocumentsByFields(page, resultsPerPage int, fields []entity.Field) []IndexResponse {

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
		return []IndexResponse{}
	}
	defer res.Body.Close()

	var (
		response []IndexResponse
		body     indexHits
	)

	if err := json.NewDecoder(res.Body).Decode(&body.Hits); err != nil {
		fmt.Println(err.Error())
		return []IndexResponse{}
	}

	for _, v := range body.Hits {
		response = append(response, v.Source)
	}

	return response
}

// func (esi *ElasticsearchIndexer) getDocById(modelColor string) IndexResponse {

// 	req := esapi.GetRequest{
// 		Index:      esi.IndexName,
// 		DocumentID: modelColor,
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	res, err := req.Do(ctx, esi.ElasticClient)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return IndexResponse{}
// 	}
// 	fmt.Println(res)
// 	defer res.Body.Close()

// 	var (
// 		response IndexResponse
// 		body     document
// 	)
// 	body.Source = &response

// 	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
// 		return IndexResponse{}
// 	}

// 	return response
// }
