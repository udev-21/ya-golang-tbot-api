package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewRestrictChatMember(userID int64, permissions types.ChatPermissions) *RestrictChatMember {
	return &RestrictChatMember{
		UserID:      userID,
		Permissions: permissions,
	}
}

// https://core.telegram.org/bots/api#restrictchatmember
type RestrictChatMember struct {
	UserID      int64                 `json:"user_id"`
	Permissions types.ChatPermissions `json:"revoke_messages"`
	UntilDate   *int64                `json:"until_date,omitempty"`
}

func (bcm *RestrictChatMember) Endpoint() string {
	return "RestrictChatMember"
}

func (bcm *RestrictChatMember) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(bcm)
}

func (bcm *RestrictChatMember) WithUntilDate(untilDate int64) {
	bcm.UntilDate = &untilDate
}
