package models

type Link struct {
    Rel  string `json:"rel"`
    Href string `json:"href"`
}

// Product representa um produto com detalhes completos incluindo avaliações e links associados.
type Product struct {
    NewId           string `json:"new_id"`
    OldId           string `json:"id"`
    Name            string `json:"name"`
    AverageRate     string `json:"average_rate"`
    TotalRecommends int    `json:"total_recommends"`
    TotalOpinions   int    `json:"total_opinions"`
    Links           []Link `json:"links"`
}
