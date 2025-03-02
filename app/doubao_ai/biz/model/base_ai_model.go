package model

type BaseRequest struct {
	UserId   int32
	Question string
	Uuid     string
	Scenario Scenario
}

type BaseResponse struct {
	StatusCode int32
	StatusMsg  string
	Data       interface{}
}

type AIConversationHandler interface {
	GetPrompt() string
	ProcessResponse(content string) (interface{}, error)
}
