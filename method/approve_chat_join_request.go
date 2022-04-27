package method

import (
	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

func NewApproveChatJoinRequest(userID int64) *ApproveChatJoinRequest {
	return &ApproveChatJoinRequest{
		UserID: userID,
	}
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
type ApproveChatJoinRequest struct {
	UserID int64 `json:"user_id"`
}

func (acjr *ApproveChatJoinRequest) Endpoint() string {
	return "approveChatJoinRequest"
}

func (acjr *ApproveChatJoinRequest) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(acjr)
}
