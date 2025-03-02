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
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	ID           int64  `json:"id,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
	CategoryID   int64  `json:"category_id,omitempty"`
}

type ProductSearchMapping struct {
	Mappings Mappings `json:"mappings,omitempty"`
}
type Mappings struct {
	Properties Properties `json:"properties,omitempty"`
}

type Properties struct {
	Name         *Field `json:"name,omitempty"`
	Description  *Field `json:"description,omitempty"`
	ID           *Field `json:"id,omitempty"`
	CategoryName *Field `json:"category_name,omitempty"`
	CategoryID   *Field `json:"category_id,omitempty"`
}

type Field struct {
	Type     string `json:"type,omitempty"`
	Analyzer string `json:"analyzer,omitempty"`
	Index    bool   `json:"index,omitempty"`
}

var ProductSearchMappingSetting = ProductSearchMapping{
	Mappings: Mappings{
		Properties: Properties{
			Name: &Field{
				Type:     "text",
				Analyzer: "ik_smart",
				Index:    true,
			},
			Description: &Field{
				Type:     "text",
				Analyzer: "ik_smart",
				Index:    true,
			},
			ID: &Field{
				Type:  "keyword",
				Index: true,
			},
			CategoryName: &Field{
				Type:     "text",
				Analyzer: "ik_smart",
				Index:    true,
			},
			CategoryID: &Field{
				Type:  "keyword",
				Index: true,
			},
		},
	},
}

type ProductSearchDoc struct {
	Name        interface{} `json:"name,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

type ProductSearchSource []string

type ProductSearchQueryBody struct {
	Query  *ProductSearchQuery  `json:"query,omitempty"`
	Doc    *ProductSearchDoc    `json:"doc,omitempty"`
	Source *ProductSearchSource `json:"_source,omitempty"`
	From   *int64               `json:"from,omitempty"`
	Size   *int64               `json:"size,omitempty"`
}
type ProductSearchTermQuery map[string]interface{}

type ProductSearchBoolQuery struct {
	Must   *[]*ProductSearchQuery `json:"must,omitempty"`
	Should *[]*ProductSearchQuery `json:"should,omitempty"`
}

type ProductSearchQuery struct {
	MultiMatch *ProductSearchMultiMatchQuery `json:"multi_match,omitempty"`
	Match      *ProductSearchMatchQuery      `json:"match,omitempty"`
	MatchAll   *All                          `json:"match_all,omitempty"`
	Term       *ProductSearchTermQuery       `json:"term,omitempty"`
	Bool       *ProductSearchBoolQuery       `json:"bool,omitempty"`
}
type All struct {
}
type ProductSearchMultiMatchQuery struct {
	Query  string   `json:"query,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

type ProductSearchMatchQuery struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProductBulkUpdate struct {
	Update ProductBulkBody `json:"update,omitempty"`
}
type ProductBulkBody struct {
	DocID interface{} `json:"_id,omitempty"`
}

type ProductBulkDoc struct {
	Doc ProductSearchDoc `json:"doc,omitempty"`
}
