package services

import (
    "github.com/go-rod/rod"
    "log"
    "WebScrapingGo/models"
)

type Scraper struct{}

func NewScraper() *Scraper {
    return &Scraper{}
}

func (s *Scraper) FetchProductCode(url string) (*models.Product, error) {
    browser := rod.New().MustConnect()
    defer browser.MustClose()

    page := browser.MustPage(url)

    // Espera até que o elemento esteja pronto
    el := page.MustWaitLoad().MustElement("[data-trustvox-product-code]")

    // Obtém o valor do atributo data-trustvox-product-code
    productCode, err := el.Attribute("data-trustvox-product-code")
    if err != nil {
        log.Printf("Erro ao obter o atributo data-trustvox-product-code: %v", err)
        return nil, err
    }

    return &models.Product{NewId: *productCode}, nil
}
