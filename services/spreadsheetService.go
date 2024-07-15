package services

import (
	"fmt"
	"log"
    "errors"
	"github.com/tealeg/xlsx"
	"WebScrapingGo/models"
)

type SpreadsheetService struct{}

// NewSpreadsheetService cria uma nova instância de SpreadsheetService
func NewSpreadsheetService() *SpreadsheetService {
    return &SpreadsheetService{}
}

// ReadSheet lê os dados do produto de um arquivo Excel
func (ss *SpreadsheetService) ReadSheet(filename string) ([]models.SheetData, error) {
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

// UpdateSheet atualiza a coluna "Ação" no arquivo Excel para o registro na linha rowIndex
func (ss *SpreadsheetService) UpdateSheet(filename string, rowIndex int, action string) error {
	// Abre o arquivo Excel para escrita
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Verifica se há pelo menos uma planilha no arquivo
	if len(wb.Sheets) == 0 {
		return errors.New("no sheets found in the workbook")
	}

	// Acessa a primeira planilha do arquivo
	sheet := wb.Sheets[0]

	// Verifica se a linha especificada está dentro do intervalo
	if rowIndex >= len(sheet.Rows) || rowIndex < 1 {
		return fmt.Errorf("row index %d out of range", rowIndex)
	}

	// Acessa a linha especificada na planilha (considerando que rowIndex começa de 0)
	row := sheet.Rows[rowIndex]

	// Ajusta o número de células para garantir que haja pelo menos 5 células na linha
	for len(row.Cells) <= 4 {
		row.AddCell()
	}

	// Define o valor da célula "Ação" na coluna 5 (índice 4)
	row.Cells[4].SetValue(action)

	// Salva as alterações de volta no arquivo Excel
	if err := wb.Save(filename); err != nil {
		return fmt.Errorf("failed to save workbook: %w", err)
	}

	return nil
}
