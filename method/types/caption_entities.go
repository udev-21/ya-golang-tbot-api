package types

import "github.com/udev-21/ya-golang-tbot-api/types"

type CaptionEntitier struct {
	CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
}

func (ce *CaptionEntitier) WithCaptionEntities(c []types.MessageEntity) {
	ce.CaptionEntities = c
}
