package main

import (
	"fmt"

	"github.com/dnogueir-org/seeker/internal/services"
	"github.com/dnogueir-org/seeker/repository/indexers"
	"github.com/dnogueir-org/seeker/web/server"
)

func main() {
	fmt.Println("Server starting")

	indexer := indexers.ElasticsearchIndexer{}
	navigationService := services.NavigationService{
		Indexer: &indexer,
	}
	webServer := server.MakeNewWebServer(navigationService)
	webServer.Serve()
}
