package controllers

import (
	"WebScrapingGo/models"
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

func (pc *ScraperController) FetchAndDisplayProduct(url string, oldProduct *models.OldProduct) (*models.Product, error) {
	product, err := pc.ScraperService.FetchProductCode(url)
	if err != nil {
		views.DisplayError(err)
		return nil, err
	}
	views.DisplayProduct(product, oldProduct)
	return product, nil
}
