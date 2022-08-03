package indexers

import "github.com/dnogueir-org/seeker/internal/entity"

type Indexer interface {
	Navigate(page, resultsPerPage int, sorting, scoringProfile string, fields []entity.Field) IndexResponse
}

type document struct {
	Source interface{} `json:"_source"`
}

type IndexResponse struct {
	Name                string  `json:"name"`
	Nickname            string  `json:"nickname"`
	Description         string  `json:"description"`
	ShortDescription    string  `json:"shortDescription"`
	ReleaseDateMembers  string  `json:"releaseDateMembers"`
	ExclusiveProduct    bool    `json:"exclusiveProduct"`
	Sneaker             bool    `json:"sneaker"`
	DisplayWithoutStock bool    `json:"displayWithoutStock"`
	GroupName           string  `json:"groupName"`
	SubGroupName        string  `json:"subGroupName"`
	CategoryName        string  `json:"categoryName"`
	ModalityName        string  `json:"modalityName"`
	GenderName          string  `json:"genderName"`
	BrandName           string  `json:"brandName"`
	LineName            string  `json:"lineName"`
	AgeGroupName        string  `json:"ageGroupName"`
	ReleaseDate         string  `json:"releaseDate"`
	URL                 string  `json:"url"`
	ModelColor          string  `json:"modelColor"`
	BasePrice           float64 `json:"basePrice"`
	CostPrice           float64 `json:"costPrice"`
	PromotionalPrice    float64 `json:"promotionalPrice"`
	Discount            int     `json:"discount"`
	ColorName           string  `json:"colorName"`
	ColorID             string  `json:"colorId"`
	HasStock            bool    `json:"hasStock"`
	Status              string  `json:"status"`
	Skus                []struct {
		Ean       string `json:"ean"`
		SizeName  string `json:"sizeName"`
		Available bool   `json:"available"`
		StyleCode string `json:"styleCode"`
		Sku       string `json:"sku"`
	} `json:"skus"`
	Filters struct {
		Caimento struct {
			Value []string `json:"value"`
		} `json:"caimento"`
		Season struct {
			Value []string `json:"value"`
		} `json:"season"`
		Tecidos struct {
			Value []string `json:"value"`
		} `json:"tecidos"`
		Tipodecamiseta struct {
			Value []string `json:"value"`
		} `json:"tipodecamiseta"`
		Tendencia struct {
			Value []string `json:"value"`
		} `json:"tendencia"`
		Cor struct {
			Value []string `json:"value"`
		} `json:"cor"`
		Snkrs struct {
			Value []bool `json:"value"`
		} `json:"snkrs"`
		Categorias struct {
			Value []string `json:"value"`
		} `json:"categorias"`
		Esportes struct {
			Value []string `json:"value"`
		} `json:"esportes"`
		Datadelancamento struct {
			Value []string `json:"value"`
		} `json:"datadelancamento"`
		Tamanho struct {
			Value []string `json:"value"`
		} `json:"tamanho"`
		Emestoque struct {
			Value []bool `json:"value"`
		} `json:"emestoque"`
		Genero struct {
			Value []string `json:"value"`
		} `json:"genero"`
		Marca struct {
			Value []string `json:"value"`
		} `json:"marca"`
		Modelos struct {
			Value []string `json:"value"`
		} `json:"modelos"`
		Idade struct {
			Value []string `json:"value"`
		} `json:"idade"`
		Tipodeproduto struct {
			Value []string `json:"value"`
		} `json:"tipodeproduto"`
		Preco struct {
			Value []string `json:"value"`
		} `json:"preco"`
		Ofertas struct {
			Value []string `json:"value"`
		} `json:"ofertas"`
	} `json:"filters"`
	Faceting []struct {
		Name  string   `json:"name"`
		Value []string `json:"value"`
	} `json:"faceting"`
	HistoricalData struct {
		LogCart     float64 `json:"logCart"`
		LogPurchase float64 `json:"logPurchase"`
		LogPdp      float64 `json:"logPdp"`
	} `json:"historicalData"`
}
