package golangtbotapi

import (
	"os"

	"github.com/udev21/golang-tbot-api/methods"
	"github.com/udev21/golang-tbot-api/types"
)

func (b *BotAPI) NewMessage(text string) *methods.Message {
	return methods.NewSendMessage(text)
}

func (b *BotAPI) NewInputFileTelegram(fileID string) *types.InputFileTelegram {
	return &types.InputFileTelegram{
		FileID: fileID,
	}
}

func (b *BotAPI) NewInputFileLocal(file *os.File) *types.InputFileLocal {
	result, err := types.NewInputFileLocal(file)
	if err != nil {
		if b.debug {
			b.logger.Fatal(err.Error())
		}
	}
	return result
}

// must be in http format
func (b *BotAPI) NewInputFileUrl(url string) *types.InputFileUrl {
	return &types.InputFileUrl{
		Url: url,
	}
}
