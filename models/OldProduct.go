package models

type OldLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

// OldProduct representa um produto antigo com detalhes completos incluindo avaliações e links associados.
type OldProduct struct {
	OldId              string `json:"id"`
	OldName            string `json:"name"`
	OldAverageRate     string `json:"average_rate"`
	OldTotalRecommends int    `json:"total_recommends"`
	OldTotalOpinions   int    `json:"total_opinions"`
	OldLinks           []OldLink `json:"old_links"`
}
