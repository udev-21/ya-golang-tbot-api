package types

import (
	"fmt"
	"reflect"
)

type BotCommandScope struct {
	Type   string      `json:"type"`
	ChatID interface{} `json:"chat_id,omitempty"`
	UserID *int64      `json:"user_id,omitempty"`
}

const (
	BotCommandScopeDefaultType               = "default"
	BotCommandScopeAllPrivateChatsType       = "all_private_chats"
	BotCommandScopeAllGroupChatsType         = "all_group_chats"
	BotCommandScopeAllChatAdministratorsType = "all_chat_administrators"
	BotCommandScopeChatType                  = "chat"
	BotCommandScopeChatAdministratorsType    = "chat_administrators"
	BotCommandScopeChatMemberType            = "chat_member"
)

func NewBotCommandScopeDefault() BotCommandScope {
	return BotCommandScope{
		Type: BotCommandScopeDefaultType,
	}
}

func NewBotCommandScopeAllPrivateChats() BotCommandScope {
	return BotCommandScope{
		Type: BotCommandScopeAllPrivateChatsType,
	}
}

func NewBotCommandScopeAllGroupChats() BotCommandScope {
	return BotCommandScope{
		Type: BotCommandScopeAllPrivateChatsType,
	}
}

// chatID Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
// must be int64 or string
func NewBotCommandScopeChat(chatID interface{}) (BotCommandScope, error) {
	if isValidChatID(chatID) {
		return newBotCommandScope(BotCommandScopeChatType, chatID), nil
	}
	return BotCommandScope{}, fmt.Errorf("chatID must be string or int64")
}

// chatID Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
// must be int64 or string
func NewBotCommandScopeChatAdministrators(chatID interface{}) (BotCommandScope, error) {
	if isValidChatID(chatID) {
		return newBotCommandScope(BotCommandScopeChatAdministratorsType, chatID), nil
	}
	return BotCommandScope{}, fmt.Errorf("chatID must be string or int64")
}

// chatID Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
// must be int64 or string
func NewBotCommandScopeChatMember(userID int64, chatID interface{}) (BotCommandScope, error) {
	if isValidChatID(chatID) {
		return BotCommandScope{
			Type:   BotCommandScopeChatMemberType,
			UserID: &userID,
			ChatID: chatID,
		}, nil
	}
	return BotCommandScope{}, fmt.Errorf("chatID must be string or int64")
}

func newBotCommandScope(scopeType string, chatID interface{}) BotCommandScope {
	return BotCommandScope{
		Type:   scopeType,
		ChatID: chatID,
	}
}

func isValidChatID(chatID interface{}) bool {
	kind := reflect.TypeOf(chatID).Kind()
	return kind == reflect.String || kind == reflect.Int64
}
