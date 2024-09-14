package prepare

type SourceJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Attributes  []struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	} `json:"attributes"`
}

type WhitelistCSV struct {
	Address   string `csv:"address" json:"address"`
	OG        string `csv:"og" json:"og"`
	Whitelist string `csv:"whitelist" json:"whitelist"`
}
