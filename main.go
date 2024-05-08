package main

import (
	"fmt"
	"os"
	"WebScrapingGo/controllers"
	"WebScrapingGo/services"
	"WebScrapingGo/views"
)

func main() {
	    // Check if the necessary command line arguments are provided
		if len(os.Args) < 4 {
			fmt.Println("Uso: <executável> <storeID> <productID> <storeToken>")
			return
		}
	
		// Extract command line arguments
		storeID := os.Args[1]
		productID := os.Args[2]
		storeToken := os.Args[3]
	
		// Retrieve the product details from the Trustvox API
		product, err := services.RetrieveOldProduct(storeID, productID, storeToken)
		if err != nil {
			views.DisplayError(err)
			return
		}
	
	// Inicialização do serviço de scraping e do controller
	scraperService := services.NewScraper()
	controller := controllers.NewProductController(scraperService)
	controller.FetchAndDisplayProduct("https://www.canecadodev.com/caneca-go-lang-commands-cheat-sheet-preta")
	//controller.FetchAndDisplayProduct("https://sleepcalm.com.br/produto/colchao-plus")

	// Display the retrieved product details
	views.DisplayProduct(product)
}