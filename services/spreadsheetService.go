package services

import (
	"log"
	"github.com/tealeg/xlsx"
	"WebScrapingGo/models"
)

// ReadSheet reads product data from an Excel file
func ReadSheet(filename string) ([]models.SheetData, error) {
	var sheetDataList []models.SheetData
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	sheet := wb.Sheets[0]
	for i, row := range sheet.Rows {
		if i == 0 {
			continue // Ignora a primeira linha (cabeçalho)
		}
		if len(row.Cells) >= 4 {
			oldProductName := row.Cells[0].Value
			oldProductCode := row.Cells[1].Value
			newProductCode := row.Cells[2].Value
			productURL := row.Cells[3].Value

			// Verifica se a linha está vazia em todas as colunas
			if oldProductName == "" && oldProductCode == "" && newProductCode == "" && productURL == "" {
				log.Println("Linha vazia encontrada, encerrando o programa.")
				break
			}

			log.Printf("Read row - OldProductName: %s, OldProductCode: %s, NewProductCode: %s, ProductURL: %s",
				oldProductName, oldProductCode, newProductCode, productURL) // Log row data

			sheetDataList = append(sheetDataList, models.SheetData{
				OldProductName: oldProductName,
				OldProductCode: oldProductCode,
				NewProductCode: newProductCode,
				ProductURL:     productURL,
			})
		}
	}
	return sheetDataList, nil
}

// UpdateSheet updates the "Ação" column in the Excel file
func UpdateSheet(filename string, rowIndex int, action string) error {
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}

	sheet := wb.Sheets[0]

	// Adiciona a célula "Ação" se necessário
	if rowIndex+1 < len(sheet.Rows) {
		row := sheet.Rows[rowIndex+1] // +1 para ajustar a linha ignorada
		for len(row.Cells) <= 4 {     // A coluna "Ação" é a quinta (índice 4)
			row.AddCell()
		}
		row.Cells[4].Value = action // Define o valor da célula "Ação"
	}

	return wb.Save(filename)
}
