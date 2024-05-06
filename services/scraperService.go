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
    log.Println("Conectando ao navegador...")
    browser := rod.New().MustConnect()
    defer browser.MustClose()

    log.Printf("Abrindo página: %s\n", url)
    page := browser.MustPage(url)

    log.Println("Esperando até que um dos elementos esteja pronto...")
    // Espera até que um dos elementos esteja pronto
    el, err := page.MustWaitLoad().Element("[data-trustvox-product-code], [data-trustvox-product-code-js]")
    if err != nil {
        log.Printf("Erro ao esperar pelo elemento: %v", err)
        return nil, err
    }

    var productCode string

    log.Println("Obtendo o valor do atributo data-trustvox-product-code...")
    attr, err := el.Attribute("data-trustvox-product-code")
    if err == nil && attr != nil {
        productCode = *attr
    } else {
        log.Println("Não foi possível obter o atributo data-trustvox-product-code, tentando data-trustvox-product-code-js...")
        attr, err = el.Attribute("data-trustvox-product-code-js")
        if err == nil && attr != nil {
            productCode = *attr
        } else {
            log.Printf("Erro ao obter o atributo data-trustvox-product-code-js: %v", err)
            return nil, err
        }
    }

    log.Println("Produto encontrado com sucesso!")
    return &models.Product{NewId: productCode}, nil
}
