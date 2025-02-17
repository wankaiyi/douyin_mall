package value

type ValueModel struct {
	//创建时间
	CreatedAt int64 `json:"created_at"`
	//本地缓存时间 单位毫秒
	Duration int64 `json:"duration"`
	//数据值
	Value any `json:"value"`
}

func (model *ValueModel) GetDefaultValue() any {
	if model.Value == nil {
		return nil
	}
	return model.Value
}
