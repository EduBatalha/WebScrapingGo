package models

import (
	"context"
	"cloud.google.com/go/language/apiv1"
	languagepb "cloud.google.com/go/language/apiv1/languagepb"
)

type TextAnalyzer struct {
	client *language.Client
}

func NewTextAnalyzer() (*TextAnalyzer, error) {
	client, err := language.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	return &TextAnalyzer{client: client}, nil
}

func (ta *TextAnalyzer) AnalyzeText(text string) (*languagepb.AnalyzeEntitiesResponse, error) {
	req := &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	}
	return ta.client.AnalyzeEntities(context.Background(), req)
}

func (ta *TextAnalyzer) CompareTexts(apiText, scrapingText string) (bool, error) {
	apiEntities, err := ta.AnalyzeText(apiText)
	if err != nil {
		return false, err
	}

	scrapingEntities, err := ta.AnalyzeText(scrapingText)
	if err != nil {
		return false, err
	}

	apiEntityMap := make(map[string]bool)
	for _, entity := range apiEntities.Entities {
		apiEntityMap[entity.Name] = true
	}

	var matchCount int
	for _, entity := range scrapingEntities.Entities {
		if _, found := apiEntityMap[entity.Name]; found {
			matchCount++
		}
	}

	// Verifica se a similaridade é maior ou igual a 70%
	similarity := float64(matchCount) / float64(len(apiEntities.Entities))
	return similarity >= 0.7, nil
}
