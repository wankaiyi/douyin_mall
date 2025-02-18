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

// NewValueModel 创建ValueModel,duration本地缓存时间,单位毫秒,value为数据值
func NewValueModel(duration int64, value any) *ValueModel {
	return &ValueModel{
		Duration: duration,
		Value:    value,
	}
}
