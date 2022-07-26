package indexers

import "github.com/dnogueir-org/seeker/internal/entity"

type Indexer interface {
	Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) IndexResponseHits
}

type IndexHits struct {
	Hits interface{} `json:"hits"`
}

type IndexResponseHits struct {
	Total IndexTotalHits      `json:"total"`
	Hits  []IndexDocumentHits `json:"hits"`
}

type IndexTotalHits struct {
	Value int `json:"value"`
}

type IndexDocumentHits struct {
	Id     string   `json:"_id"`
	Source Document `json:"_source"`
}

type Document struct {
	Name               string             `json:"name"`
	Nickname           string             `json:"nickname"`
	Description        string             `json:"description"`
	ShortDescription   string             `json:"shortDescription"`
	ReleaseDateMembers string             `json:"releaseDateMembers"`
	Sneaker            bool               `json:"sneaker"`
	GroupName          string             `json:"groupName"`
	SubGroupName       string             `json:"subGroupName"`
	CategoryName       string             `json:"categoryName"`
	ModalityName       string             `json:"modalityName"`
	GenderName         string             `json:"genderName"`
	BrandName          string             `json:"brandName"`
	LineName           string             `json:"lineName"`
	AgeGroupName       string             `json:"ageGroupName"`
	ReleaseDate        string             `json:"releaseDate"`
	URL                string             `json:"url"`
	ModelColor         string             `json:"modelColor"`
	BasePrice          float64            `json:"basePrice"`
	PromotionalPrice   float64            `json:"promotionalPrice"`
	Discount           float64            `json:"discount"`
	ColorName          string             `json:"colorName"`
	ColorID            string             `json:"colorId"`
	HasStock           bool               `json:"hasStock"`
	Status             string             `json:"status"`
	Skus               []IndexResponseSku `json:"skus"`
}

type IndexResponseSku struct {
	Ean       string `json:"ean"`
	SizeName  string `json:"sizeName"`
	Available bool   `json:"available"`
	StyleCode string `json:"styleCode"`
	Sku       string `json:"sku"`
}
