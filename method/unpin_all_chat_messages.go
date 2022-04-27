package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
)

func NewUnpinAllChatMessages() *UnpinAllChatMessages {
	return &UnpinAllChatMessages{}
}

type UnpinAllChatMessages struct{}

func (upcm *UnpinAllChatMessages) Endpoint() string {
	return "unpinAllChatMessages"
}

func (uacm *UnpinAllChatMessages) Params() (myTypes.Params, error) {
	return make(myTypes.Params), nil
}
