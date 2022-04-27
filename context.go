package yagolangtbotapi

import (
	"io"

	"github.com/udev-21/ya-golang-tbot-api/method"
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

type Context interface {
	// Bot returns the bot instance.
	Bot() *BotAPI

	Send(to interface{}, paylad myTypes.Sendable) (*types.ApiResponse, error)

	// Update returns the original update.
	Update() types.Update

	// Message - New incoming message of any kind — text, photo, sticker, etc.
	Message() *types.Message

	// EditedMessage - New version of a message that is known to the bot and was edited
	EditedMessage() *types.Message

	// ChannelPost - New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost() *types.Message

	// ChannelPostEdited - New version of a channel post that is known to the bot and was edited
	ChannelPostEdited() *types.Message

	// InlineQuery - New incoming inline query
	InlineQuery() *types.InlineQuery

	// InlineQuery - New incoming inline query
	ChoosenInlineResult() *types.ChosenInlineResult

	// Callback returns stored callback if such presented.
	CallbackQuery() *types.CallbackQuery

	// ShippingQuery returns stored shipping query if such presented.
	ShippingQuery() *types.ShippingQuery

	// PreCheckoutQuery returns stored pre checkout query if such presented.
	PreCheckoutQuery() *types.PreCheckoutQuery

	// Poll returns stored poll if such presented.
	Poll() *types.Poll

	// PollAnswer returns stored poll answer if such presented.
	PollAnswer() *types.PollAnswer

	// MyChatMember - The bot's chat member status was updated in a chat.
	// For private chats, this update is received only when the bot is blocked or unblocked by the user.
	MyChatMember() *types.ChatMemberUpdated

	// ChatMember - A chat member's status was updated in a chat.
	// The bot must be an administrator in the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	ChatMember() *types.ChatMemberUpdated

	// ChatJoinRequest
	ChatJoinRequest() *types.ChatJoinRequest

	// User returns the current recipient, depending on the context type.
	// Returns nil if user is not presented.
	Sender() *types.User

	// Chat returns the current chat, depending on the context type.
	// Returns nil if chat is not presented.
	Chat() *types.Chat

	// after this line goes "bot embedded" methods like
	GetFile(fileID string) (*types.File, error)
	DownloadFile(file *types.File, dest io.Writer) error
	LeaveChat(chat *types.Chat) error
	GetChat(chatID string) (*types.Chat, error)
	GetChatAdministrators(chat *types.Chat) (types.ChatMembers, error)
	DeleteChatPhoto(chat *types.Chat) error
	CreateChatInviteLink(chat *types.Chat, content *method.CreateChatInviteLink) (*types.ChatInviteLink, error)
	EditChatInviteLink(chat *types.Chat, content *method.EditChatInviteLink) (*types.ChatInviteLink, error)
	RevokeChatInviteLink(chat *types.Chat, content *method.RevokeChatInviteLink) (*types.ChatInviteLink, error)
	DeleteMessage(chat *types.Chat, messageID int64) error
	SetChatStickerSet(chat *types.Chat, stickerSetName string) error
	DeleteChatStickerSet(chat *types.Chat) error
	GetChatMember(chat *types.Chat, userID int64) (*types.ChatMember, error)
	GetChatMemberCount(chat *types.Chat) (*int64, error)
	GetStickerSet(name string) (*types.StickerSet, error)
	ExportChatInviteLink(chat *types.Chat) (*string, error)
	SetChatAdministratorCustomTitle(chat *types.Chat, userID int64, title string) error
}
