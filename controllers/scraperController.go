package controllers

import (
	"WebScrapingGo/services"
	"WebScrapingGo/models"
	"WebScrapingGo/views"
)

type ScraperController struct {
	ScraperService *services.Scraper
}

func NewProductController(scraper *services.Scraper) *ScraperController {
	return &ScraperController{
		ScraperService: scraper,
	}
}

func (pc *ScraperController) FetchAndDisplayProduct(url string, oldProduct *models.OldProduct) {
	product, err := pc.ScraperService.FetchProductCode(url)
	if err != nil {
		views.DisplayError(err)
		return
	}
	views.DisplayProduct(product, oldProduct)
}
