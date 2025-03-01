package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type Scenario int8

const (
	OrderInquiry   Scenario = iota + 1 // 1 - 查询订单对话消息
	MockPlaceOrder                     // 2 - 模拟下单对话消息
)

type Message struct {
	Base
	UserId   int32    `gorm:"not null;type:int;comment:'用户ID'"`
	Role     Role     `gorm:"type:varchar(64);not null;comment:'角色：user-用户，assistant-AI'"`
	Content  string   `gorm:"not null; type:text;comment:'消息内容'"`
	Uuid     string   `gorm:"type:varchar(64);not null;index;comment:'一次会话的唯一标识'"`
	Scenario Scenario `gorm:"type:tinyint;not null;comment:'对话场景，1-查询订单，2-模拟下单'"`
}

func (m Message) TableName() string {
	return "tb_message"
}

func CreateMessage(db *gorm.DB, ctx context.Context, message *Message) error {
	result := db.Create(message)
	return result.Error
}

func ConversionExist(db *gorm.DB, ctx context.Context, userId int32, uuid string) (bool, error) {
	var count int64
	result := db.WithContext(ctx).Model(&Message{}).
		Where(&Message{
			Uuid:   uuid,
			UserId: userId,
		}).
		Count(&count)
	if result.Error == nil {
		return count > 0, nil
	}
	return false, result.Error
}

func GetChatHistoryByUuid(db *gorm.DB, ctx context.Context, uuid string) ([]Message, error) {
	var messages []Message
	result := db.WithContext(ctx).
		Model(&Message{}).
		Where(&Message{Uuid: uuid}).
		Order("id asc").
		Find(&messages)
	return messages, result.Error
}

func (m Message) String() string {
	return fmt.Sprintf("Message [ID: %d, UserID: %d, Role: %s, Content: %s, UUID: %s, Scenario: %d, CreatedAt: %s, UpdatedAt: %s]",
		m.ID, m.UserId, string(m.Role), m.Content, m.Uuid, m.Scenario, m.CreatedAt, m.UpdatedAt)
}
