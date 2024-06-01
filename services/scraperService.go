package services

import (
	"WebScrapingGo/models"
	"log"
    "fmt"
	"github.com/go-rod/rod"
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
	el, err := page.MustWaitLoad().Element("[data-trustvox-product-code], [data-trustvox-product-code-js]")
	if err != nil {
		log.Printf("Erro ao esperar pelo elemento: %v", err)
		return nil, err
	}

	productCode, err := getAttributeValue(el, "data-trustvox-product-code", "data-trustvox-product-code-js")
	if err != nil {
		log.Printf("Erro ao obter atributos: %v", err)
		return nil, err
	}

	log.Println("Produto encontrado com sucesso!")
	return &models.Product{NewId: productCode}, nil
}

func getAttributeValue(el *rod.Element, attrs ...string) (string, error) {
	for _, attr := range attrs {
		value, err := el.Attribute(attr)
		if err == nil && value != nil {
			return *value, nil
		}
	}
	return "", fmt.Errorf("nenhum dos atributos encontrados: %v", attrs)
}
