package services

import (
	"github.com/dnogueir-org/seeker/internal/converters"
	"github.com/dnogueir-org/seeker/repository/indexers"
)

type NavigationRequest struct {
	BusinessUnit   string
	Page           int
	ResultsPerPage int
	RestrictSearch bool
	Sorting        string
	ScoringProfile string
	Fields         []string
}

type NavigationService struct {
	Indexer indexers.Indexer
}

func (ns *NavigationService) Navigate(request NavigationRequest) (interface{}, error) {
	filterFields := converters.ConvertFilterField(request.Fields)
	indexerResponse := ns.Indexer.Navigate(request.Page, request.ResultsPerPage, request.Sorting, request.ScoringProfile, filterFields)

	return indexerResponse, nil
}
