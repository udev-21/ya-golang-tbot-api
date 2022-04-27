package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendChatAction(action string) *ChatAction {
	return &ChatAction{
		Action: action,
	}
}

type ChatAction struct {
	Action string `json:"action"`
}

const (
	ChatActionTyping          = "typing"
	ChatActionUploadPhoto     = "upload_photo"
	ChatActionUploadVideo     = "upload_video"
	ChatActionRecordVideo     = "record_video"
	ChatActionUploadVideoNote = "upload_video_note"
	ChatActionRecordVideoNote = "record_video_note"
	ChatActionUploadVoice     = "upload_voice"
	ChatActionRecordVoice     = "record_voice"
	ChatActionUploadDocument  = "upload_document"
	ChatActionChooseSticker   = "choose_sticker"
	ChatActionFindLocation    = "find_location"
)

func (sca *ChatAction) Endpoint() string {
	return "sendChatAction"
}

func (sca *ChatAction) Params() (myTypes.Params, error) {
	return utils.ConvertToMapStringInterface(sca)
}
