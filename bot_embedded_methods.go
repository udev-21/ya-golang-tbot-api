package yagolangtbotapi

import (
	"encoding/json"
	"io"
	"log"

	"github.com/udev-21/ya-golang-tbot-api/method"
	"github.com/udev-21/ya-golang-tbot-api/types"
)

func (ba *BotAPI) GetMe() (*types.User, error) {
	res, err := ba.request("getMe", nil)
	if err != nil {
		return nil, newError(err.Error())
	}
	var response types.User
	bytes, err := json.Marshal(res.Result)
	if err != nil {
		return nil, newError(err.Error())
	}
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, newError(err.Error())
	}
	return &response, nil
}

func (ba *BotAPI) LogOut() (*types.ApiResponse, error) {
	return ba.request("logOut", nil)
}

func (ba *BotAPI) Close() (*types.ApiResponse, error) {
	return ba.request("close", nil)
}

func (ba *BotAPI) GetFile(fileID string) (*types.File, error) {
	res, err := ba.request("getFile", map[string]interface{}{
		"file_id": fileID,
	})
	if err != nil {
		return nil, newError(err.Error())
	}
	if !res.OK {
		msg, err := json.Marshal(res)
		if err != nil {
			return nil, newError(err.Error())
		}
		return nil, newError(string(msg))
	}
	var response types.File
	bytes, err := json.Marshal(res.Result)
	if err != nil {
		return nil, newError(err.Error())
	}
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, newError(err.Error())
	}
	return &response, nil
}

func (ba *BotAPI) DownloadFile(file *types.File, dest io.Writer) error {
	if file == nil {
		msg := "file is required"
		writeLog(LogLevelError, ba.logger, msg)
		return newError(msg)
	}
	if dest == nil {
		msg := "destination is required"
		writeLog(LogLevelError, ba.logger, msg)
		return newError(msg)
	}
	if file.FilePath == nil {
		msg := "filePath is required"
		writeLog(LogLevelError, ba.logger, msg)
		return newError(msg)
	}
	res, err := ba.httpClient.Get(ba.getFilePath(*file.FilePath))
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return err
	}
	defer res.Body.Close()
	// if not binary then just return error
	if res.Header.Get("Content-Type") == "application/json" {
		return newError("file not found")
	}
	_, err = io.Copy(dest, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func (ba *BotAPI) LeaveChat(chat *types.Chat) error {
	chatID, err := getChatID(chat)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return newError("chat is required")
	}

	res, err := ba.request("leaveChat", map[string]interface{}{
		"chat_id": chatID,
	})

	if err != nil {
		return newError(err.Error())
	}

	if !res.OK {
		return newError("not leaved")
	}
	return nil
}

func (ba *BotAPI) GetChat(chatID string) (*types.Chat, error) {
	if len(chatID) == 0 {
		writeLog(LogLevelError, ba.logger, "chat is required")
		return nil, newError("chat is required")
	}

	res, err := ba.request("getChat", map[string]interface{}{
		"chat_id": chatID,
	})
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}

	var response types.Chat
	bytes, err := json.Marshal(res.Result)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}
	if err = json.Unmarshal(bytes, &response); err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}
	return &response, nil
}

func (ba *BotAPI) GetChatAdministrators(chat *types.Chat) (types.ChatMembers, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}

	res, err := request(ba.getPath("getChatAdministrators"), map[string]interface{}{
		"chat_id": chatID,
	}, ba.httpClient)

	if err != nil {
		return nil, err
	}
	log.Println(string(res))
	var response struct {
		types.ApiResponse
		ChatMembers types.ChatMembers `json:"result,omitempty"`
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}

	if !response.OK {
		writeLog(LogLevelError, ba.logger, "something went wrong")
		return nil, newError("something went wrong")
	}
	return response.ChatMembers, nil
}

func (ba *BotAPI) DeleteChatPhoto(chat *types.Chat) error {
	chatID, err := getChatID(chat)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return newError(err.Error())
	}

	res, err := request(ba.getPath("deleteChatPhoto"), map[string]interface{}{
		"chat_id": chatID,
	}, ba.httpClient)

	if err != nil {
		return err
	}
	log.Println(string(res))
	var response types.ApiResponse

	err = json.Unmarshal(res, &response)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return newError(err.Error())
	}

	if !response.OK {
		writeLog(LogLevelError, ba.logger, "something went wrong")
		return newError("something went wrong")
	}
	return nil
}

func (ba *BotAPI) CreateChatInviteLink(chat *types.Chat, content *method.CreateChatInviteLink) (*types.ChatInviteLink, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}
	params, err := content.Params()
	if err != nil {
		return nil, err
	}
	params["chat_id"] = chatID
	res, err := request(ba.getPath("createChatInviteLink"), params, ba.httpClient)
	if err != nil {
		return nil, err
	}
	log.Println(string(res))
	var response struct {
		OK   bool                 `json:"ok"`
		Link types.ChatInviteLink `json:"result"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}

	if !response.OK {
		writeLog(LogLevelError, ba.logger, "something went wrong")
		return nil, newError("something went wrong")
	}
	return &response.Link, nil
}

func (ba *BotAPI) EditChatInviteLink(chat *types.Chat, content *method.EditChatInviteLink) (*types.ChatInviteLink, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}
	params, err := content.Params()
	if err != nil {
		return nil, err
	}
	params["chat_id"] = chatID
	res, err := request(ba.getPath("editChatInviteLink"), params, ba.httpClient)
	if err != nil {
		return nil, err
	}
	log.Println(string(res))
	var response struct {
		OK   bool                 `json:"ok"`
		Link types.ChatInviteLink `json:"result"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}

	if !response.OK {
		writeLog(LogLevelError, ba.logger, "something went wrong")
		return nil, newError("something went wrong")
	}
	return &response.Link, nil
}

func (ba *BotAPI) RevokeChatInviteLink(chat *types.Chat, content *method.RevokeChatInviteLink) (*types.ChatInviteLink, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}
	params, err := content.Params()
	if err != nil {
		return nil, err
	}
	params["chat_id"] = chatID
	res, err := request(ba.getPath("editChatInviteLink"), params, ba.httpClient)
	if err != nil {
		return nil, err
	}
	log.Println(string(res))
	var response struct {
		OK   bool                 `json:"ok"`
		Link types.ChatInviteLink `json:"result"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		writeLog(LogLevelError, ba.logger, err.Error())
		return nil, newError(err.Error())
	}

	if !response.OK {
		writeLog(LogLevelError, ba.logger, "something went wrong")
		return nil, newError("something went wrong")
	}
	return &response.Link, nil
}

func (ba *BotAPI) DeleteMessage(chat *types.Chat, messageID int64) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}
	res, err := ba.request("deleteMessage", map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	})
	if err != nil {
		return err
	} else if !res.OK {
		return newError("not deleted")
	}
	var tres bool
	json.Unmarshal(res.Result, &tres)
	if !tres {
		return newError("not deleted")
	}
	return nil
}

func (ba *BotAPI) SetChatStickerSet(chat *types.Chat, stickerSetName string) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}
	res, err := ba.request("setChatStickerSet", map[string]interface{}{
		"chat_id":          chatID,
		"sticker_set_name": stickerSetName,
	})
	if err != nil {
		return err
	} else if !res.OK {
		return newError("not set")
	}
	var tres bool
	json.Unmarshal(res.Result, &tres)
	if !tres {
		return newError("not set")
	}
	return nil
}

func (ba *BotAPI) DeleteChatStickerSet(chat *types.Chat) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}
	res, err := ba.request("deleteChatStickerSet", map[string]interface{}{
		"chat_id": chatID,
	})
	if err != nil {
		return err
	} else if !res.OK {
		return newError("not deleted")
	}
	var tres bool
	json.Unmarshal(res.Result, &tres)
	if !tres {
		return newError("not deleted")
	}
	return nil
}

func (ba *BotAPI) GetChatMember(chat *types.Chat, userID int64) (*types.ChatMember, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		return nil, err
	}
	res, err := ba.request("getChatMember", map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	} else if !res.OK {
		return nil, newError("getChatMember failed")
	}
	var tres types.ChatMember
	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return nil, err
	}

	return &tres, nil
}

func (ba *BotAPI) GetChatMemberCount(chat *types.Chat) (*int64, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		return nil, err
	}
	res, err := ba.request("getChatMemberCount", map[string]interface{}{
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	} else if !res.OK {
		return nil, newError("getChatMemberCount failed")
	}
	var tres int64

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return nil, newError("getChatMemberCount unmarshal error: " + err.Error())
	}

	return &tres, nil
}

func (ba *BotAPI) GetStickerSet(name string) (*types.StickerSet, error) {
	res, err := ba.request("getStickerSet", map[string]interface{}{
		"name": name,
	})

	if err != nil {
		return nil, err
	} else if !res.OK {
		return nil, newError("getStickerSet failed")
	}
	var tres types.StickerSet

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return nil, newError("getStickerSet unmarshal error: " + err.Error())
	}

	return &tres, nil
}

func (ba *BotAPI) ExportChatInviteLink(chat *types.Chat) (*string, error) {
	chatID, err := getChatID(chat)
	if err != nil {
		return nil, err
	}

	res, err := ba.request("exportChatInviteLink", map[string]interface{}{
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	} else if !res.OK {
		return nil, newError("exportChatInviteLink failed")
	}
	var tres string

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return nil, newError("exportChatInviteLink unmarshal error: " + err.Error())
	}

	return &tres, nil
}

func (ba *BotAPI) SetChatAdministratorCustomTitle(chat *types.Chat, userID int64, title string) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}

	res, err := ba.request("setChatAdministratorCustomTitle", map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
		"title":   title,
	})

	if err != nil {
		return err
	} else if !res.OK {
		return newError("setChatAdministratorCustomTitle failed")
	}
	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("setChatAdministratorCustomTitle unmarshal error: " + err.Error())
	}
	if !tres {
		return newError("setChatAdministratorCustomTitle failed")
	}

	return nil
}

func (ba *BotAPI) BanChatSenderChat(chat *types.Chat, senderChatID int64) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}

	res, err := ba.request("banChatSenderChat", map[string]interface{}{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})

	if err != nil {
		return err
	} else if !res.OK {
		return newError("BanChatSenderChat failed")
	}
	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("BanChatSenderChat unmarshal error: " + err.Error())
	}
	if !tres {
		return newError("BanChatSenderChat failed")
	}

	return nil
}

func (ba *BotAPI) UnBanChatSenderChat(chat *types.Chat, senderChatID int64) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}

	res, err := ba.request("unbanChatSenderChat", map[string]interface{}{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})

	if err != nil {
		return err
	} else if !res.OK {
		return newError("unbanChatSenderChat failed")
	}
	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("unbanChatSenderChat unmarshal error: " + err.Error())
	}
	if !tres {
		return newError("unbanChatSenderChat failed")
	}

	return nil
}

func (ba *BotAPI) SetChatPermissions(chat *types.Chat, permissions *types.ChatPermissions) error {
	chatID, err := getChatID(chat)
	if err != nil {
		return err
	}

	if permissions == nil {
		return newError("permissions required")
	}

	res, err := ba.request("setChatPermissions", map[string]interface{}{
		"chat_id":     chatID,
		"permissions": permissions,
	})

	if err != nil {
		return err
	} else if !res.OK {
		return newError("setChatPermissions failed")
	}
	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("setChatPermissions unmarshal error: " + err.Error())
	}
	if !tres {
		return newError("setChatPermissions failed")
	}

	return nil
}

func (ba *BotAPI) SetMyCommands(payload *method.SetMyCommands) error {
	if payload == nil {
		return newError("payload required")
	}

	res, err := ba.Send(nil, payload)

	if err != nil {
		return err
	} else if !res.OK {
		return newError("SetMyCommands failed")
	}

	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("SetMyCommands unmarshal error: " + err.Error())
	}

	if !tres {
		return newError("SetMyCommands failed")
	}

	return nil
}

func (ba *BotAPI) DeleteMyCommands(payload *method.DeleteMyCommands) error {
	if payload == nil {
		return newError("payload required")
	}

	res, err := ba.Send(nil, payload)

	if err != nil {
		return err
	} else if !res.OK {
		return newError("DeleteMyCommands failed")
	}

	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("DeleteMyCommands unmarshal error: " + err.Error())
	}

	if !tres {
		return newError("DeleteMyCommands failed")
	}

	return nil
}

func (ba *BotAPI) GetMyCommands(payload *method.GetMyCommands) ([]types.BotCommand, error) {
	if payload == nil {
		return nil, newError("payload required")
	}

	res, err := ba.Send(nil, payload)

	if err != nil {
		return nil, err
	} else if !res.OK {
		return nil, newError("GetMyCommands failed")
	}

	var tres []types.BotCommand

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return nil, newError("GetMyCommands unmarshal error: " + err.Error())
	}

	return tres, nil
}

func (ba *BotAPI) SetChatMenuButton(button *method.SetChatMenuButton) error {
	if button == nil {
		return newError("button required")
	}

	res, err := ba.Send(nil, button)

	if err != nil {
		return err
	} else if !res.OK {
		return newError("SetChatMenuButton failed")
	}

	var tres bool

	if err := json.Unmarshal(res.Result, &tres); err != nil {
		return newError("SetChatMenuButton unmarshal error: " + err.Error())
	} else if !tres {
		return newError("SetChatMenuButton failed")
	}

	return nil
}
