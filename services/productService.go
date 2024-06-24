package services

import (
	"WebScrapingGo/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ProductService interface {
	RetrieveOldProduct(storeID, productID, storeToken string) (*models.OldProduct, error)
}

type productService struct {
	client *http.Client
}

func NewProductService(client *http.Client) ProductService {
	return &productService{client: client}
}

func (ps *productService) RetrieveOldProduct(storeID, productID, storeToken string) (*models.OldProduct, error) {
	url := fmt.Sprintf("https://trustvox.com.br/api/stores/%s/products/%s", storeID, productID)
	log.Printf("URL: %s", url) // Log URL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.trustvox.com; version=1")
	req.Header.Add("Authorization", "Bearer "+storeToken)
	log.Printf("Request headers: %v", req.Header) // Log headers

	res, err := ps.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	defer res.Body.Close()

	log.Printf("Response status: %d", res.StatusCode) // Log status code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	log.Printf("Response body: %s", string(body)) // Log response body
	var data models.OldProduct
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &data, nil
}
