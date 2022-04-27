package golangtbotapi

import (
	"encoding/json"
	"io"
	"log"

	"github.com/udev-21/golang-tbot-api/method"
	"github.com/udev-21/golang-tbot-api/types"
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
