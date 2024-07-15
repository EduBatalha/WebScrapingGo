package main

import (
	"WebScrapingGo/controllers"
	"WebScrapingGo/models"
	"WebScrapingGo/services"
	"WebScrapingGo/views"
	"fmt"
	"net/http"
	"os"
	"time"
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

	// Defina a variável de ambiente GOOGLE_APPLICATION_CREDENTIALS
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/caminho/para/arquivo.json")

	// Inicializa os serviços necessários
	client := &http.Client{Timeout: 10 * time.Second}
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Inicializa os serviços específicos
	productService := services.NewProductService(client)
	scraperService := services.NewScraperService(browser)
	sheetService := services.NewSpreadsheetService()

	// Inicializa a console view
	consoleView := views.ConsoleView{}

	// Inicializa o TextAnalyzer
	textAnalyzer, err := models.NewTextAnalyzer()
	if err != nil {
		consoleView.DisplayError(err)
		return
	}

	// Inicializa o controlador de produtos
	productController := controllers.NewProductController(productService, scraperService, *sheetService, consoleView, textAnalyzer)

	// Executa o processo de leitura da planilha e processamento de produtos
	if err := processProducts(storeID, storeToken, filename, productController); err != nil {
		consoleView.DisplayError(err) // Exibe o erro usando ConsoleView
	}
}

func processProducts(storeID, storeToken, filename string, productController *controllers.ProductController) error {
	// Processa os produtos usando o controller
	return productController.ProcessProducts(storeID, storeToken, filename)
}
