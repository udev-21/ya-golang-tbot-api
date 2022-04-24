package method

import (
	"fmt"

	myTypes "github.com/udev-21/golang-tbot-api/method/types"
	"github.com/udev-21/golang-tbot-api/utils"
)

type MediaGroup struct {
	Media []myTypes.InputMedia `json:"media"`

	myTypes.DisableNotificationer
	myTypes.ProtectContenter
	myTypes.ReplyToMessager
}

func (mg *MediaGroup) Endpoint() string {
	return "sendMediaGroup"
}

func (mg *MediaGroup) Files() []myTypes.InputFile {
	return prepareInputMediaFiles(mg.Media)
}

func (mg *MediaGroup) Params() (myTypes.Params, error) {
	params, err := utils.ConvertToMapStringInterface(mg)
	if err != nil {
		return nil, err
	}
	tmp := prepareInputMediaForParams(mg.Media)
	params["media"] = tmp
	return params, nil
}

func prepareInputMediaFiles(inputMedias []myTypes.InputMedia) []myTypes.InputFile {
	var newFiles []myTypes.InputFile

	for idx, media := range inputMedias {
		// log.Printf("\n%+v\n", media)
		if files := prepareInputMediaFile(media, idx); len(files) > 0 {
			newFiles = append(newFiles, files...)
		}
	}

	return newFiles
}
func prepareInputMediaFile(inputMedia myTypes.InputMedia, idx int) []myTypes.InputFile {
	var files []myTypes.InputFile

	switch m := inputMedia.(type) {
	case *myTypes.InputMediaPhoto:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("attach://file_%d", idx))
			files = append(files, tmp)
		} else {
			files = append(files, m.Media)
		}
	case *myTypes.InputMediaVideo:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("attach://file_%d", idx))
			files = append(files, tmp)
		} else {
			files = append(files, m.Media)
		}

		if m.Thumb != nil {
			if tmp, ok := (*m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetAttachName(fmt.Sprintf("attach://file_%d_thumb", idx))
				files = append(files, tmp)
			} else {
				files = append(files, *m.Thumb)
			}
		}
	case *myTypes.InputMediaDocument:

		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("attach://file_%d", idx))
			files = append(files, tmp)
		} else {
			files = append(files, m.Media)
		}

		if m.Thumb != nil {
			if tmp, ok := (*m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetAttachName(fmt.Sprintf("attach://file_%d_thumb", idx))
				files = append(files, tmp)
			} else {
				files = append(files, *m.Thumb)
			}
		}
	case *myTypes.InputMediaAudio:

		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("attach://file_%d", idx))
			files = append(files, tmp)
		} else {
			files = append(files, m.Media)
		}

		if m.Thumb != nil {
			if tmp, ok := (*m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetAttachName(fmt.Sprintf("attach://file_%d_thumb", idx))
				files = append(files, tmp)
			} else {
				files = append(files, *m.Thumb)
			}
		}

	}
	return files
}
func prepareInputMediaForParams(inputMedia []myTypes.InputMedia) []myTypes.InputMedia {
	newMedia := make([]myTypes.InputMedia, len(inputMedia))
	copy(newMedia, inputMedia)

	for idx, media := range inputMedia {
		if param := prepareInputMediaParam(media, idx); param != nil {
			newMedia[idx] = param
		}
	}

	return newMedia
}

// It is expected to be used in conjunction with prepareInputMediaFile.
func prepareInputMediaParam(inputMedia myTypes.InputMedia, idx int) myTypes.InputMedia {

	switch m := inputMedia.(type) {
	case *myTypes.InputMediaPhoto:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("file_%d", idx))
			m.Media = tmp
		}
		return m
	case *myTypes.InputMediaVideo:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("file_%d", idx))
			m.Media = tmp
		}

		if m.Thumb != nil {
			if tmp, ok := (*m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetAttachName(fmt.Sprintf("file_%d_thumb", idx))
				*m.Thumb = tmp
			}
		}

		return m
	case *myTypes.InputMediaAudio:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("file_%d", idx))
			m.Media = tmp
		}

		if m.Thumb != nil {
			if tmp, ok := (*m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetAttachName(fmt.Sprintf("file_%d_thumb", idx))
				*m.Thumb = tmp
			}
		}

		return m
	case *myTypes.InputMediaDocument:
		if tmp, ok := m.Media.(myTypes.Uploadable); ok {
			tmp.SetAttachName(fmt.Sprintf("file_%d", idx))
			m.Media = tmp
		}

		if m.Thumb != nil {
			if tmp, ok := (*m.Thumb).(myTypes.Uploadable); ok {
				tmp.SetAttachName(fmt.Sprintf("file_%d_thumb", idx))
				*m.Thumb = tmp
			}
		}
		return m
	}

	return nil
}
