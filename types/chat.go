package types

// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID                    int64            `json:"id"`
	Type                  string           `json:"type"`
	Title                 *string          `json:"title,omitempty"`
	Username              *string          `json:"username,omitempty"`
	FirstName             *string          `json:"first_name,omitempty"`
	LastName              *string          `json:"last_name,omitempty"`
	Photo                 *ChatPhoto       `json:"photo,omitempty"`
	BIO                   *string          `json:"bio,omitempty"`
	HasPrivateForwards    *bool            `json:"has_private_forwards,omitempty"`
	Description           *string          `json:"description,omitempty"`
	InviteLink            *string          `json:"invite_link,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
	Permissions           *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay         *int64           `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime *int64           `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent   *bool            `json:"has_protected_content,omitempty"`
	StickerSetName        *string          `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      *bool            `json:"can_set_sticker_set,omitempty"`
	LinkedChatID          *int64           `json:"linked_chat_id,omitempty"`
	Location              *ChatLocation    `json:"location,omitempty"`
}

const (
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSuperGroup = "supergroup"
	ChatTypeChannel    = "channel"
)

func (ch Chat) IsPrivate() bool {
	return ch.Type == ChatTypePrivate
}

func (ch Chat) IsGroup() bool {
	return ch.Type == ChatTypeGroup
}

func (ch Chat) IsSupergroup() bool {
	return ch.Type == ChatTypeSuperGroup
}

func (ch Chat) IsChannel() bool {
	return ch.Type == ChatTypeChannel
}

func (ch Chat) IsGroupOrSupergroup() bool {
	return ch.IsGroup() || ch.IsSupergroup()
}
