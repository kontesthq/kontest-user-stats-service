package codechef

type CodeChefUser struct {
	Success       bool              `json:"success"`
	Profile       string            `json:"profile"`
	Name          string            `json:"name"`
	CurrentRating int               `json:"currentRating"`
	HighestRating int               `json:"highestRating"`
	CountryFlag   string            `json:"countryFlag"`
	CountryName   string            `json:"countryName"`
	GlobalRank    int               `json:"globalRank"`
	CountryRank   int               `json:"countryRank"`
	Stars         string            `json:"stars"`
	HeatMap       []HeatMapEntry    `json:"heatMap"`
	RatingData    []RatingDataEntry `json:"ratingData"`
}
