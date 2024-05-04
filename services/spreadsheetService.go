// spreadsheet_service.go
package services

import (
	"github.com/tealeg/xlsx"
	"WebScrapingGo/models"
)

func ReadSheet(filename string) ([]models.URLData, error) {
	var urls []models.URLData
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	sheet := wb.Sheets[0]
	for _, row := range sheet.Rows {
		cell := row.Cells[0]
		url := cell.String()
		urls = append(urls, models.URLData{URL: url})
	}
	return urls, nil

}