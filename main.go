package main

import (
	"fmt"
	"os"
	"WebScrapingGo/controllers"
	"WebScrapingGo/services"
	"WebScrapingGo/views"
)

func main() {
	// Verifica se os argumentos necessários da linha de comando foram fornecidos
	if len(os.Args) < 4 {
		fmt.Println("Uso: <executável> <storeID> <productID> <storeToken>")
		return
	}

	// Extrai os argumentos da linha de comando
	storeID := os.Args[1]
	productID := os.Args[2]
	storeToken := os.Args[3]

	// Recupera os detalhes do produto antigo da API Trustvox
    oldProduct, err := services.RetrieveOldProduct(storeID, productID, storeToken)
    if err != nil {
        views.DisplayError(err)
        return
    }
	
	// Inicializa o serviço de scraping e o controlador
	scraperService := services.NewScraper()
	controller := controllers.NewProductController(scraperService)

	// URL passada para FetchAndDisplayProduct
	productURL := "https://sleepcalm.com.br/produto/colchao-plus?gad_source=1"
	controller.FetchAndDisplayProduct(productURL, oldProduct)

	// Exibe os detalhes do produto recuperado
	views.DisplayProduct(nil, oldProduct)
}
