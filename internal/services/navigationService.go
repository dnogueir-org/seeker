package services

import (
	"fmt"
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

type NavigationService struct{}

func (ns *NavigationService) Navigate(request NavigationRequest) (string, error) {
	for _, v := range request.Fields {
		fmt.Println(v)
	}
	return "teste", nil
}
