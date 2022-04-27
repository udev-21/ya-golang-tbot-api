package golangtbotapi

import (
	"io"

	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/types"
)

type nativeContext struct {
	bot    *BotAPI
	update types.Update
}

var _ Context = (*nativeContext)(nil)

func (n *nativeContext) Bot() *BotAPI {
	return n.bot
}

func (n *nativeContext) Update() types.Update {
	return n.update
}

func (n *nativeContext) Message() *types.Message {
	return n.update.Message
}

func (n *nativeContext) EditedMessage() *types.Message {
	return n.update.EditedMessage
}

func (n *nativeContext) ChannelPost() *types.Message {
	return n.update.ChannelPost
}

func (n *nativeContext) ChannelPostEdited() *types.Message {
	return n.update.EditedChannelPost
}

func (n *nativeContext) InlineQuery() *types.InlineQuery {
	return n.update.InlineQuery
}

func (n *nativeContext) ChoosenInlineResult() *types.ChosenInlineResult {
	return n.update.ChosenInlineResult
}

func (n *nativeContext) CallbackQuery() *types.CallbackQuery {
	return n.update.CallbackQuery
}

func (n *nativeContext) ShippingQuery() *types.ShippingQuery {
	return n.update.ShippingQuery
}

func (n *nativeContext) PreCheckoutQuery() *types.PreCheckoutQuery {
	return n.update.PreCheckoutQuery
}

func (n *nativeContext) Poll() *types.Poll {
	return n.update.Poll
}

func (n *nativeContext) PollAnswer() *types.PollAnswer {
	return n.update.PollAnswer
}

func (n *nativeContext) MyChatMember() *types.ChatMemberUpdated {
	return n.update.MyChatMember
}

func (n *nativeContext) ChatMember() *types.ChatMemberUpdated {
	return n.update.ChatMember
}

func (n *nativeContext) ChatJoinRequest() *types.ChatJoinRequest {
	return n.update.ChatJoinRequest
}

func (n *nativeContext) Sender() *types.User {
	switch {
	case n.update.CallbackQuery != nil:
		return &n.update.CallbackQuery.Sender

	case n.update.InlineQuery != nil:
		return &n.update.InlineQuery.Sender
	case n.update.ChosenInlineResult != nil:
		return &n.update.ChosenInlineResult.Sender
	case n.update.ShippingQuery != nil:
		return &n.update.ShippingQuery.Sender
	case n.update.PreCheckoutQuery != nil:
		return &n.update.PreCheckoutQuery.Sender
	case n.update.PollAnswer != nil:
		return &n.update.PollAnswer.Sender
	case n.update.MyChatMember != nil:
		return &n.update.MyChatMember.Sender
	case n.update.ChatMember != nil:
		return &n.update.ChatMember.Sender
	case n.update.ChatJoinRequest != nil:
		return &n.update.ChatJoinRequest.Sender
	case n.update.Message != nil:
		return n.update.Message.Sender
	case n.update.CallbackQuery != nil:
		return &n.update.CallbackQuery.Sender
	case n.update.EditedMessage != nil:
		return n.update.EditedMessage.Sender
	case n.update.ChannelPost != nil:
		if n.update.ChannelPost.PinnedMessage != nil {
			return n.update.ChannelPost.PinnedMessage.Sender
		}
		return n.update.ChannelPost.Sender
	case n.update.EditedChannelPost != nil:
		return n.update.EditedChannelPost.Sender
	default:
		return nil
	}
}

func (n *nativeContext) Chat() *types.Chat {
	switch {
	case n.update.Message != nil:
		return &n.update.Message.Chat
	case n.update.CallbackQuery != nil:
		return &n.update.CallbackQuery.Message.Chat
	case n.update.EditedMessage != nil:
		return &n.update.EditedMessage.Chat
	case n.update.ChannelPost != nil:
		if n.update.ChannelPost.PinnedMessage != nil {
			return &n.update.ChannelPost.PinnedMessage.Chat
		}
		return &n.update.ChannelPost.Chat
	case n.update.EditedChannelPost != nil:
		return &n.update.EditedChannelPost.Chat
	default:
		return nil
	}
}

func (n *nativeContext) Send(to interface{}, payload myTypes.Sendable) (*types.ApiResponse, error) {
	return n.bot.Send(to, payload)
}

func (n *nativeContext) GetFile(fileID string) (*types.File, error) {
	return n.bot.GetFile(fileID)
}

func (n *nativeContext) DownloadFile(file *types.File, dest io.Writer) error {
	return n.bot.DownloadFile(file, dest)
}

func (n *nativeContext) LeaveChat(chat *types.Chat) error {
	return n.bot.LeaveChat(chat)
}

func (n *nativeContext) GetChat(chatID string) (*types.Chat, error) {
	return n.bot.GetChat(chatID)
}

func (n *nativeContext) GetChatAdministrators(chat *types.Chat) (types.ChatMembers, error) {
	return n.bot.GetChatAdministrators(chat)
}
