package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
)

func NewUnpinAllChatMessages() *UnpinAllChatMessages {
	return &UnpinAllChatMessages{}
}

type UnpinAllChatMessages struct{}

func (upcm *UnpinAllChatMessages) Endpoint() string {
	return "unpinChatMessage"
}

func (uacm *UnpinAllChatMessages) Params() (myTypes.Params, error) {
	return make(myTypes.Params), nil
}
