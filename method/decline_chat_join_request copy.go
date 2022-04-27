package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewDeclineChatJoinRequest(userID int64) *DeclineChatJoinRequest {
	return &DeclineChatJoinRequest{
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#declinechatjoinrequest
type DeclineChatJoinRequest struct {
	UserID int64 `json:"user_id"`
}

func (dcjr *DeclineChatJoinRequest) Endpoint() string {
	return "declineChatJoinRequest"
}

func (dcjr *DeclineChatJoinRequest) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(dcjr)
}
