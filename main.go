package main

import (
	"WebScrapingGo/controllers"
	"WebScrapingGo/services"
)

func main() {
	scraperService := services.NewScraper()
	controller := controllers.NewProductController(scraperService)
	controller.FetchAndDisplayProduct("https://sleepcalm.com.br/produto/colchao-plus")
}
