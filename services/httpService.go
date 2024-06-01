// httpService.go
package services

import (
	"WebScrapingGo/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RetrieveOldProduct(storeID, productID, storeToken string) (*models.OldProduct, error) {
	url := fmt.Sprintf("https://trustvox.com.br/api/stores/%s/products/%s", storeID, productID)
	client := &http.Client{
		Timeout: 10 * time.Second, // Define um timeout para o cliente HTTP
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.trustvox.com; version=1")
	req.Header.Add("Authorization", "Bearer "+storeToken)

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error performing request: %v", err)
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %d", res.StatusCode)
		return nil, fmt.Errorf("received non-OK HTTP status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, fmt.Errorf("error reading response body: %w", err)
	}


	var data models.OldProduct
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}


	return &data, nil
}
