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
