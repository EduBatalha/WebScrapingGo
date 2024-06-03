package services

import (
	"context"
	"log"
	"time"
	"fmt"
	"WebScrapingGo/models"
	"github.com/go-rod/rod"
)

type Scraper struct {
	browser *rod.Browser
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func NewScraperWithBrowser(browser *rod.Browser) *Scraper {
	return &Scraper{browser: browser}
}

func (s *Scraper) FetchProductCode(url string) (*models.Product, error) {
	log.Println("Abrindo página:", url)
	page := s.browser.MustPage(url)

	// Cria um contexto com timeout de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	var el *rod.Element
	var err error

	go func() {
		defer close(done)
		log.Println("Esperando até que um dos elementos esteja pronto...")
		el, err = page.Element("[data-trustvox-product-code], [data-trustvox-product-code-js]")
	}()

	select {
	case <-ctx.Done():
		log.Println("Tempo limite de 10 segundos excedido.")
		return nil, ctx.Err()
	case <-done:
		if err != nil {
			log.Printf("Erro ao esperar pelo elemento: %v", err)
			return nil, err
		}
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
