package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewBanChatMember(userID int64) *BanChatMember {
	return &BanChatMember{
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#banchatmember
type BanChatMember struct {
	UserID         int64  `json:"user_id"`
	UntilDate      *int64 `json:"until_date,omitempty"`
	RevokeMessages *bool  `json:"revoke_messages,omitempty"`
}

func (bcm *BanChatMember) Endpoint() string {
	return "banChatMember"
}

func (bcm *BanChatMember) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(bcm)
}

func (bcm *BanChatMember) WithUntilDate(untilDate int64) {
	bcm.UntilDate = &untilDate
}

func (bcm *BanChatMember) WithRevokeMessages(untilDate int64) {
	bcm.RevokeMessages = new(bool)
	*bcm.RevokeMessages = true
}
