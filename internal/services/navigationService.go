package services

import (
	"fmt"
)

type NavigationService struct{}

func (ns *NavigationService) Navigate(filterPairs []string) string {
	for _, v := range filterPairs {
		fmt.Println(v)
	}
	return "teste"
}
