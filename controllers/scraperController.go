package controllers

import (
	"WebScrapingGo/models"
	"WebScrapingGo/services"
)

type ScraperController struct {
	ScraperService services.ScraperService
}

func NewScraperController(scraperService services.ScraperService) *ScraperController {
	return &ScraperController{ScraperService: scraperService}
}

func (sc *ScraperController) FetchProductCode(url string) (*models.Product, error) {
	return sc.ScraperService.FetchProductCode(url)
}
