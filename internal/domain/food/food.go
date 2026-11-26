package food

// Food represents a catalog item and its nutrition data.
type Food struct {
	ID               string    `json:"id"`
	SourceID         string    `json:"source_id"`
	CanonicalName    string    `json:"canonical_name"`
	AlternativeNames []string  `json:"alternative_names"`
	Category         string    `json:"category"`
	SourceGroup      string    `json:"source_group"`
	ServingSize      *string   `json:"serving_size"`
	BaseAmount       float64   `json:"base_amount"`
	Tags             []string  `json:"tags"`
	AllergenInfo     []string  `json:"allergen_info"`
	Nutrients        Nutrients `json:"nutrients"`
	SourceStatus     string    `json:"source_status"`
	Names            Names     `json:"names"`
}

// Names maps locale identifiers to localized food names.
type Names map[string]string

// Nutrients maps stable nutrient keys to values reported by each source.
type Nutrients map[string]NutrientValue

// NutrientValue contains a nutrient's display label, unit, and source values.
// A nil source value represents missing data; a pointer to zero is a measured zero.
type NutrientValue struct {
	Label  string   `json:"label"`
	Unit   string   `json:"unit"`
	TACO   *float64 `json:"taco"`
	USDA   *float64 `json:"usda"`
	Outros *float64 `json:"outros"`
}

// FoodSearch contains the filters and pagination for a catalog search.
type FoodSearch struct {
	Query    string
	Category string
	Limit    int
	Offset   int
}

// FoodPage is a page of matching foods and the total number of matches.
type FoodPage struct {
	Items  []Food
	Total  int
	Limit  int
	Offset int
}
