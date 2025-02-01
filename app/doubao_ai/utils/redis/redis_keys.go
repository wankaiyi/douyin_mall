package redis

import "fmt"

const (
	chatHistoryKey = "doubao_ai:chat_history:%s"
)

func GetChatHistoryKey(uuid string) string {
	return fmt.Sprintf(chatHistoryKey, uuid)
}
