package vo

type ProductSearchAllDataVo struct {
	Took     int64 `json:"took"`
	TimedOut bool  `json:"timed_out"`
	Shards   struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	} `json:"_shards"`
	MaxScore float64 `json:"max_score"`
	Hits     struct {
		Total struct {
			Value    int64  `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []struct {
			ID     string              `json:"_id"`
			Score  float64             `json:"_score"`
			Index  string              `json:"_index"`
			Type   string              `json:"_type"`
			Source ProductSearchDataVo `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
type ProductSearchDataVo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductSearchMapping struct {
	Mappings Mappings `json:"mappings"`
}
type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Name        Field `json:"name"`
	Description Field `json:"description"`
}

type Field struct {
	Type     string `json:"type"`
	Analyzer string `json:"analyzer"`
}

var ProductSearchMappingSetting = ProductSearchMapping{
	Mappings: Mappings{
		Properties: Properties{
			Name: Field{
				Type:     "text",
				Analyzer: "ik_smart",
			},
			Description: Field{
				Type:     "text",
				Analyzer: "ik_smart",
			},
		},
	},
}
