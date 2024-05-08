package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "WebScrapingGo/models"
)

func RetrieveOldProduct(storeID, productID, storeToken string) (*models.Product, error) {
    url := fmt.Sprintf("https://trustvox.com.br/api/stores/%s/products/%s", storeID, productID)
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Printf("Error creating request: %v", err)
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/vnd.trustvox.com; version=1")
    req.Header.Add("Authorization", "Bearer " + storeToken)

    res, err := client.Do(req)
    if err != nil {
        log.Printf("Error performing request: %v", err)
        return nil, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Printf("Error reading response body: %v", err)
        return nil, err
    }

    var data models.Product
    if err := json.Unmarshal(body, &data); err != nil {
        log.Printf("Error parsing JSON: %v", err)
        return nil, err
    }

    return &data, nil
}
