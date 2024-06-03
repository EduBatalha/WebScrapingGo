package main

import (
	"WebScrapingGo/controllers"
	"WebScrapingGo/services"
	"WebScrapingGo/views"
	"context"
	"fmt"
	"os"

	"github.com/go-rod/rod"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: <executable> <storeID> <storeToken> <filename>")
		return
	}

	storeID := os.Args[1]
	storeToken := os.Args[2]
	filename := os.Args[3]

	// Lê os dados da planilha
	productData, err := services.ReadSheet(filename)
	if err != nil {
		views.DisplayError(err)
		return
	}

	// Inicializa o serviço de scraping e o controlador
	browser := rod.New().MustConnect()
	defer browser.MustClose()
	scraperService := services.NewScraperWithBrowser(browser)
	controller := controllers.NewProductController(scraperService)

	// Itera sobre cada produto
	for i, data := range productData {
		// Recupera os detalhes do produto antigo da API usando OldProductCode
		oldProduct, err := services.RetrieveOldProduct(storeID, data.OldProductCode, storeToken)
		if err != nil {
			views.DisplayError(err)
			services.UpdateSheet(filename, i, "ID antigo")
			continue
		}

		// Passa a URL do produto para FetchAndDisplayProduct
		product, err := controller.FetchAndDisplayProduct(data.ProductURL, oldProduct)
		if err != nil {
			if err == context.DeadlineExceeded {
				services.UpdateSheet(filename, i, "verificar")
				fmt.Println("Ocorreu um timeout. Passando para a próxima linha.")
				continue
			}
			views.DisplayError(err)
			services.UpdateSheet(filename, i, "ID novo")
			continue
		}

		// Verifica se ambos os IDs foram retornados corretamente
		if oldProduct.OldId != "" && product.NewId != "" {
			services.UpdateSheet(filename, i, "confere")
		} else if oldProduct.OldId != "" {
			services.UpdateSheet(filename, i, "ID antigo")
		} else {
			services.UpdateSheet(filename, i, "ID novo")
		}
	}
}
