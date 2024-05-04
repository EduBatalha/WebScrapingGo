package controllers

import (
	"WebScrapingGo/services"
	"WebScrapingGo/views"
)

type ScraperController struct {
	SpreadsheetPath string
}

func (sc *ScraperController) ExecuteScraping() {
	urls, err := services.ReadSheet(sc.SpreadsheetPath)
	if err != nil {
		views.ShowMessage("Erro ao ler arquivo: " + err.Error())
		return
	}

	for _, url := range urls {
		services.ScrapeURL(url.URL)
	}
}
