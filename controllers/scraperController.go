package controllers

import (
	"WebScrapingGo/services"
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

func (pc *ScraperController) FetchAndDisplayProduct(url string) {
	product, err := pc.ScraperService.FetchProductCode(url)
	if err != nil {
		views.DisplayError(err)
		return
	}
	views.DisplayProduct(product)
}
