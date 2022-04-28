package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewUnbanChatMember(userID int64) *UnbanChatMember {
	return &UnbanChatMember{
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#unbanchatmember
type UnbanChatMember struct {
	UserID       int64 `json:"user_id"`
	OnlyIfBanned *bool `json:"only_if_banned,omitempty"`
}

func (bcm *UnbanChatMember) Endpoint() string {
	return "unbanChatMember"
}

func (bcm *UnbanChatMember) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(bcm)
}

func (bcm *UnbanChatMember) WithOnlyIfBanned(untilDate int64) {
	bcm.OnlyIfBanned = new(bool)
	*bcm.OnlyIfBanned = true
}
