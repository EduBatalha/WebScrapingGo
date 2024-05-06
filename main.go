package main

import (
	"WebScrapingGo/controllers"
	"WebScrapingGo/services"
)

func main() {
	scraperService := services.NewScraper()
	controller := controllers.NewProductController(scraperService)
	//controller.FetchAndDisplayProduct("https://www.canecadodev.com/caneca-go-lang-commands-cheat-sheet-preta")
	controller.FetchAndDisplayProduct("https://sleepcalm.com.br/produto/colchao-plus")
}
