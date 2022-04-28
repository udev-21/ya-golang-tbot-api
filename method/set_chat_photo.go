package method

import (
	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSetChatPhoto(photo *myTypes.FilePath) *SetChatPhoto {
	return &SetChatPhoto{
		Photo: photo,
	}
}

// https://core.telegram.org/bots/api#setchatphoto
type SetChatPhoto struct {
	Photo myTypes.InputFile `json:"photo"`
}

func (scp *SetChatPhoto) Endpoint() string {
	return "setChatPhoto"
}

func (scp *SetChatPhoto) Params() (myTypes.Params, error) {
	params, err := utils.ConvertToMapStringInterface(scp)
	return params, err
}

func (scp *SetChatPhoto) Files() []myTypes.Uploadable {
	var files []myTypes.Uploadable
	if tmp, ok := scp.Photo.(myTypes.Uploadable); ok {
		tmp.SetField("photo")
		files = append(files, tmp)
	}
	return files
}
