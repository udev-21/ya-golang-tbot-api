package golangtbotapi

import "github.com/udev21/golang-tbot-api/types"

type Context interface {
	// Bot returns the bot instance.
	Bot() *BotAPI

	Send(to types.Chat, paylad MessagePayload) (*types.ApiResponse, error)

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
}

type nativeContext struct {
	bot    *BotAPI
	update types.Update
}

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

// implement Send function of interface Context to internalContext
func (n *nativeContext) Send(to types.Chat, payload MessagePayload) (*types.ApiResponse, error) {
	return n.bot.Send(to, payload)
}
