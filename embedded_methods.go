package golangtbotapi

import (
	"encoding/json"
	"io"

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
