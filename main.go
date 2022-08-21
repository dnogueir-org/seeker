package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/dnogueir-org/seeker/internal/services"
	"github.com/dnogueir-org/seeker/repository/indexers"
	"github.com/dnogueir-org/seeker/web/server"
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	fmt.Println("Server starting")

	cfg := elasticsearch.Config{
		Addresses: []string{"http://fisia-indexer.prd.gcp.gruposbf.com.br"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 3 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 3 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	es, _ := elasticsearch.NewClient(cfg)

	indexer := indexers.ElasticsearchIndexer{
		ElasticClient: es,
		IndexName:     "fisia-products",
	}
	navigationService := services.NavigationService{
		Indexer: &indexer,
	}
	webServer := server.MakeNewWebServer(navigationService)
	webServer.Serve()
}
