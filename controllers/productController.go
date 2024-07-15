package controllers

import (
	"fmt"
	"WebScrapingGo/models"
	"WebScrapingGo/services"
	"WebScrapingGo/views"
)

type ProductController struct {
	ProductService services.ProductService
	ScraperService services.ScraperService
	SheetService   services.SpreadsheetService
	ConsoleView    views.ConsoleView
	TextAnalyzer   *models.TextAnalyzer
}

// Função construtora para criar uma nova instância de ProductController
func NewProductController(productService services.ProductService, scraperService services.ScraperService, sheetService services.SpreadsheetService, consoleView views.ConsoleView, textAnalyzer *models.TextAnalyzer) *ProductController {
	return &ProductController{
		ProductService: productService,
		ScraperService: scraperService,
		SheetService:   sheetService,
		ConsoleView:    consoleView,
		TextAnalyzer:   textAnalyzer,
	}
}

func (pc *ProductController) ProcessProducts(storeID, storeToken, filename string) error {
	productData, err := pc.SheetService.ReadSheet(filename)
	if err != nil {
		return fmt.Errorf("error reading sheet: %w", err)
	}

	// Preenche a coluna "Ação"
	for i, data := range productData {
		rowIndex := i + 1
		if err := pc.ProcessProduct(storeID, storeToken, filename, rowIndex, data); err != nil {
			pc.ConsoleView.DisplayError(err)
		}
	}

	return nil
}

func (pc *ProductController) ProcessProduct(storeID, storeToken, filename string, rowIndex int, data models.SheetData) error {
	// Verifica se há Novo Código de Produto na planilha
	if data.NewProductCode == "" {
		err := fmt.Errorf("missing New Product Code")
		pc.ConsoleView.DisplayError(err)
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Sem ID novo")
	}

	// Recupera os detalhes do produto antigo usando API
	oldProduct, err := pc.ProductService.RetrieveOldProduct(storeID, data.OldProductCode, storeToken)
	if err != nil {
		pc.ConsoleView.DisplayError(fmt.Errorf("error retrieving old product: %w", err))
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Sem ID antigo")
	}

	// Passa a URL do produto para FetchProductDetails
	productDetails, err := pc.ScraperService.FetchProductDetails(data.ProductURL)
	if err != nil {
		pc.ConsoleView.DisplayError(fmt.Errorf("error fetching product details: %w", err))
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Erro ao buscar detalhes do produto")
	}

	// Compara os textos usando TextAnalyzer
	match, err := pc.TextAnalyzer.CompareTexts(oldProduct.OldName, productDetails.RelevantTexts)
	if err != nil {
		pc.ConsoleView.DisplayError(fmt.Errorf("error comparing texts: %w", err))
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Erro na comparação de textos")
	}

	if match {
		return pc.SheetService.UpdateSheet(filename, rowIndex, "confere")
	} else {
		return pc.SheetService.UpdateSheet(filename, rowIndex, "não confere")
	}
}
