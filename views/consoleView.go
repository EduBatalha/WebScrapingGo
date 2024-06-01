package views

import (
	"WebScrapingGo/models"
	"fmt"
)

// DisplayProduct exibe as informações do produto e do old product.
func DisplayProduct(product *models.Product, oldProduct *models.OldProduct) {
    if product != nil {
        fmt.Printf("New Product ID: %s\n", product.NewId)
        if oldProduct != nil && oldProduct.OldId != "" {
            fmt.Printf("Old Product ID: %s\n", oldProduct.OldId)
        } else {
            fmt.Println("Old Product ID não encontrado.")
        }
    }
}

// DisplayError exibe erros que ocorrem durante a execução.
func DisplayError(err error) {
	fmt.Println("Erro:", err)
}

// ShowMessage exibe uma mensagem geral.
func ShowMessage(message string) {
	fmt.Println(message)
}
