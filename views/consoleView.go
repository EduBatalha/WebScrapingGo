package views

import (
	"fmt"
	"WebScrapingGo/models"
)

// DisplayProduct exibe o novo ID do produto.
func DisplayProduct(product *models.Product) {
	if product != nil {
		fmt.Printf("New Product ID: %s\n", product.NewId)
	} else {
		fmt.Println("Product não encontrado.")
	}
}

// Retorno da API Trustvox
func ShowMessage(message string) {
    fmt.Println(message)
}

// DisplayError exibe erros que ocorrem durante a execução
func DisplayError(err error) {
	fmt.Println("Erro:", err)
}
