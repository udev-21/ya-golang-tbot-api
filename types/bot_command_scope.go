package types

type BotCommandScope interface {
	IsBotCommandScope()
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

type BaseBotCommandScope struct {
	Type string `json:"type"`
}

func (bbcs *BaseBotCommandScope) IsBotCommandScope() {}

type BotCommandScopeDefault struct {
	BaseBotCommandScope
}

func NewBotCommandScopeDefault() *BotCommandScopeDefault {
	return &BotCommandScopeDefault{
		BaseBotCommandScope: BaseBotCommandScope{
			Type: BotCommandScopeDefaultType,
		},
	}
}

type BotCommandScopeAllPrivateChats struct {
	BaseBotCommandScope
}

func NewBotCommandScopeAllPrivateChats() *BotCommandScopeAllPrivateChats {
	return &BotCommandScopeAllPrivateChats{
		BaseBotCommandScope: BaseBotCommandScope{
			Type: BotCommandScopeAllPrivateChatsType,
		},
	}
}

type BotCommandScopeAllGroupChats struct {
	BaseBotCommandScope
}

func NewBotCommandScopeAllGroupChats() *BotCommandScopeAllGroupChats {
	return &BotCommandScopeAllGroupChats{
		BaseBotCommandScope: BaseBotCommandScope{
			Type: BotCommandScopeAllPrivateChatsType,
		},
	}
}

type BotCommandScopeAllChatAdministrators struct {
	BaseBotCommandScope
}

func NewBotCommandScopeAllChatAdministrators() *BotCommandScopeAllChatAdministrators {
	return &BotCommandScopeAllChatAdministrators{
		BaseBotCommandScope: BaseBotCommandScope{
			Type: BotCommandScopeAllChatAdministratorsType,
		},
	}
}

type BotCommandScopeChat struct {
	BaseBotCommandScope
	ChatID string `json:"chat_id"`
}

func NewBotCommandScopeChat(chatID string) *BotCommandScopeChat {
	return &BotCommandScopeChat{
		BaseBotCommandScope: BaseBotCommandScope{Type: BotCommandScopeChatType},
		ChatID:              chatID,
	}
}

type BotCommandScopeChatAdministrators struct {
	BaseBotCommandScope
	ChatID string `json:"chat_id"`
}

func NewBotCommandScopeChatAdministrators(chatID string) *BotCommandScopeChatAdministrators {
	return &BotCommandScopeChatAdministrators{
		BaseBotCommandScope: BaseBotCommandScope{Type: BotCommandScopeChatAdministratorsType},
		ChatID:              chatID,
	}
}

type BotCommandScopeChatMember struct {
	BaseBotCommandScope
	ChatID string `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

func NewBotCommandScopeChatMember(chatID string, userID int64) *BotCommandScopeChatMember {
	return &BotCommandScopeChatMember{
		BaseBotCommandScope: BaseBotCommandScope{Type: BotCommandScopeChatMemberType},
		UserID:              userID,
		ChatID:              chatID,
	}
}
