package controllers

import (
	"fmt"
	"log"
	"WebScrapingGo/models"
	"WebScrapingGo/services"
	"WebScrapingGo/views"
)

type ProductController struct {
	ProductService services.ProductService
	ScraperService services.ScraperService
	SheetService   services.SpreadsheetService
	ConsoleView    views.ConsoleView
}

// Função construtora para criar uma nova instância de ProductController
func NewProductController(productService services.ProductService, scraperService services.ScraperService, sheetService services.SpreadsheetService, consoleView views.ConsoleView) *ProductController {
	return &ProductController{
		ProductService: productService,
		ScraperService: scraperService,
		SheetService:   sheetService,
		ConsoleView:    consoleView,
	}
}

func (pc *ProductController) ProcessProducts(storeID, storeToken, filename string) error {
	productData, err := pc.SheetService.ReadSheet(filename)
	if err != nil {
		return fmt.Errorf("error reading sheet: %w", err)
	}

	//Preenche a coluna "Ação" corretamente
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

	// Recupera os detalhes do produto antigo da API usando OldProductCode
	oldProduct, err := pc.ProductService.RetrieveOldProduct(storeID, data.OldProductCode, storeToken)
	if err != nil {
		pc.ConsoleView.DisplayError(fmt.Errorf("error retrieving old product: %w", err))
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Sem ID antigo")
	}

	// Passa a URL do produto para FetchProductCode
	product, err := pc.ScraperService.FetchProductCode(data.ProductURL)
	if err != nil {
		pc.ConsoleView.DisplayError(fmt.Errorf("error fetching product code: %w", err))
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Erro ao buscar código do produto")
	}

	// Adiciona prints de depuração para verificar os valores de oldProduct e product
	log.Printf("Row %d: oldProduct.OldId=%s, product.NewId=%s\n", rowIndex, oldProduct.OldId, product.NewId)

	// Verifica se ambos os IDs foram retornados corretamente
	if oldProduct.OldId != "" {
		if data.NewProductCode == product.NewId {
			log.Printf("Row %d: Both IDs match. Marking as 'confere'\n", rowIndex)
			return pc.SheetService.UpdateSheet(filename, rowIndex, "confere")
		} else {
			log.Printf("Row %d: IDs do not match. Marking as 'Sem ID novo'\n", rowIndex)
			return pc.SheetService.UpdateSheet(filename, rowIndex, "Sem ID novo")
		}
	} else {
		log.Printf("Row %d: Old ID missing. Marking as 'Sem ID antigo'\n", rowIndex)
		return pc.SheetService.UpdateSheet(filename, rowIndex, "Sem ID antigo")
	}
}
