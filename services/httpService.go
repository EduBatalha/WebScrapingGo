// http_service.go
package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"WebScrapingGo/views"
	
)

func ScrapeURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao fazer a solicitação GET: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler a resposta: %s", err)
		return
	}

	views.ShowMessage("Resposta da API: " + string(body))
};