package yagolangtbotapi

// func (b *BotAPI) NewMessage(text string) *method.Message {
// 	return method.NewSendMessage(text)
// }

// func (b *BotAPI) NewInputFileTelegram(fileID string) *types.InputFileTelegram {
// 	return &types.InputFileTelegram{
// 		FileID: fileID,
// 	}
// }

// func (b *BotAPI) NewInputFileLocal(file *os.File) *types.InputFileLocal {
// 	result, err := types.NewInputFileLocal(file)
// 	if err != nil {
// 		if b.debug {
// 			b.logger.Fatal(err.Error())
// 		}
// 	}
// 	return result
// }

// must be in http format
// func (b *BotAPI) NewInputFileUrl(url string) *types.InputFileUrl {
// 	return &types.InputFileUrl{
// 		Url: url,
// 	}
// }

func applyMiddlewares(handler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	return func(ctx Context) error {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler(ctx)
	}

}
