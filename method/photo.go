package method

import "github.com/udev-21/golang-tbot-api/types"

type Photo struct {
	ChatID                  string                `json:"chat_id"`
	Photo                   types.InputFile       `json:"photo"`
	Caption                 *string               `json:"caption,omitempty"`
	ParseMode               *string               `json:"parse_mode,omitempty"`
	CaptionEntities         []types.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification     *bool                 `json:"disable_notification,omitempty"`
	ProtectContent          *bool                 `json:"protect_content,omitempty"`
	ReplyToMessageID        *int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithutReply *bool                 `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup             interface{}           `json:"reply_markup,omitempty"`
}
