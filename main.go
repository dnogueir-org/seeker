package main

import (
	"fmt"

	"github.com/dnogueir-org/seeker/internal/services"
	"github.com/dnogueir-org/seeker/web/server"
)

func main() {
	fmt.Println("Hello World")

	navigationService := services.NavigationService{}
	webServer := server.MakeNewWebServer(navigationService)
	webServer.Serve()
}
