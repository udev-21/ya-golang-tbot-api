package method

import (
	"encoding/json"
	"os"
	"strconv"

	myTypes "github.com/udev-21/ya-golang-tbot-api/method/types"
	"github.com/udev-21/ya-golang-tbot-api/utils"
)

func NewSendMediaGroup(media []myTypes.SendMediaGroupable) *MediaGroup {
	return &MediaGroup{
		Media: media,
	}
}

// https://core.telegram.org/bots/api#sendmediagroup
type MediaGroup struct {
	Media []myTypes.SendMediaGroupable `json:"media"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
}

func (mg *MediaGroup) Endpoint() string {
	return "sendMediaGroup"
}

func (mg *MediaGroup) Files() []myTypes.Uploadable {
	return prepareInputMediaFiles(mg.Media)
}

func (mg *MediaGroup) Params() (myTypes.Params, error) {
	params, err := utils.ConvertToMapStringInterface(mg)
	if err != nil {
		return nil, err
	}
	tmp := prepareInputMediaForParams(mg.Media)
	params["media"] = tmp
	json.NewEncoder(os.Stdout).Encode(params)
	// os.Exit(1)
	return params, nil
}

func prepareInputMediaFiles(inputMedias []myTypes.SendMediaGroupable) []myTypes.Uploadable {
	var newFiles []myTypes.Uploadable

	for idx, media := range inputMedias {
		if files := prepareInputMediaFile(media, idx); len(files) > 0 {
			newFiles = append(newFiles, files...)
		}
	}
	return newFiles
}

func prepareInputMediaFile(inputMedia myTypes.SendMediaGroupable, idx int) []myTypes.Uploadable {
	var files []myTypes.Uploadable

	switch m := inputMedia.(type) {
	case *myTypes.InputMediaPhoto:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetField(tmp.CustomFileName())
			files = append(files, tmp)
		}
	case *myTypes.InputMediaVideo:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetField(tmp.CustomFileName())
			files = append(files, tmp)
		}

		if m.Thumb != nil {
			if tmp, ok := (m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetField("file_" + strconv.FormatInt(int64(idx), 10) + "_thumb")
				files = append(files, tmp)
			}
		}
	case *myTypes.InputMediaDocument:

		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetField(tmp.CustomFileName())
			files = append(files, tmp)
		}

		if m.Thumb != nil {
			if tmp, ok := (m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetField("file_" + strconv.FormatInt(int64(idx), 10) + "_thumb")
				files = append(files, tmp)
			}
		}

	case *myTypes.InputMediaAudio:

		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetField(tmp.CustomFileName())
			files = append(files, tmp)
		}

		if m.Thumb != nil {
			if tmp, ok := (m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetField("file_" + strconv.FormatInt(int64(idx), 10) + "_thumb")
				files = append(files, tmp)
			}
		}
	}
	return files
}
func prepareInputMediaForParams(inputMedia []myTypes.SendMediaGroupable) []myTypes.SendMediaGroupable {
	newMedia := make([]myTypes.SendMediaGroupable, len(inputMedia))
	copy(newMedia, inputMedia)

	for idx, media := range inputMedia {
		if param := prepareInputMediaParam(media, idx); param != nil {
			newMedia[idx] = param
		}
	}

	return newMedia
}

func prepareInputMediaParam(inputMedia myTypes.SendMediaGroupable, idx int) myTypes.SendMediaGroupable {
	switch m := inputMedia.(type) {
	case *myTypes.InputMediaPhoto:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetCustomFileName(tmp.FileName())
			m.Media = tmp
		}

		return m
	case *myTypes.InputMediaVideo:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetCustomFileName(tmp.FileName())
			m.Media = tmp
		}

		if m.Thumb != nil {
			if tmp, ok := (m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetCustomFileName("file_" + strconv.FormatInt(int64(idx), 10) + "_thumb")
				m.Thumb = tmp
			}
		}

		return m
	case *myTypes.InputMediaAudio:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetCustomFileName(tmp.FileName())
			m.Media = tmp
		}

		if m.Thumb != nil {
			if tmp, ok := (m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetCustomFileName("file_" + strconv.FormatInt(int64(idx), 10) + "_thumb")
				m.Thumb = tmp
			}
		}

		return m
	case *myTypes.InputMediaDocument:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetCustomFileName(tmp.FileName())
			m.Media = tmp
		}

		if m.Thumb != nil {
			if tmp, ok := (m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetCustomFileName("file_" + strconv.FormatInt(int64(idx), 10) + "_thumb")
				m.Thumb = tmp
			}
		}
		return m
	}

	return nil
}
