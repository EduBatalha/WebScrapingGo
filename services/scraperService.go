package services

import (
	"context"
	"fmt"
	"time"
	"WebScrapingGo/models"
	"github.com/go-rod/rod"
)

type ScraperService interface {
	FetchProductCode(url string) (*models.Product, error)
	FetchProductDetails(url string) (*models.ProductDetails, error)
}

type scraperService struct {
	browser *rod.Browser
}

func NewScraperService(browser *rod.Browser) ScraperService {
	return &scraperService{browser: browser}
}

func (ss *scraperService) FetchProductCode(url string) (*models.Product, error) {
	page := ss.browser.MustPage(url)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	var el *rod.Element
	var err error

	go func() {
		defer close(done)
		el, err = page.Element("[data-trustvox-product-code], [data-trustvox-product-code-js]")
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		if err != nil {
			return nil, err
		}
	}

	productCode, err := getAttributeValue(el, "data-trustvox-product-code", "data-trustvox-product-code-js")
	if err != nil {
		return nil, err
	}

	return &models.Product{NewId: productCode}, nil
}

func (ss *scraperService) FetchProductDetails(url string) (*models.ProductDetails, error) {
	page := ss.browser.MustPage(url)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	var el *rod.Element
	var err error

	go func() {
		defer close(done)
		el, err = page.Element("[data-trustvox-product-code], [data-trustvox-product-code-js]")
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		if err != nil {
			return nil, err
		}
	}

	productCode, err := getAttributeValue(el, "data-trustvox-product-code", "data-trustvox-product-code-js")
	if err != nil {
		return nil, err
	}

	// Extrair textos relevantes
	texts, err := page.Elements("body *")
	if err != nil {
		return nil, err
	}

	var relevantTexts string
	for _, element := range texts {
		text, err := element.Text()
		if err == nil && text != "" {
			relevantTexts += text + " "
		}
	}

	return &models.ProductDetails{
		NewId:         productCode,
		RelevantTexts: relevantTexts,
	}, nil
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
