package types

import "github.com/udev-21/golang-tbot-api/types"

type Captioner struct {
	Caption         *string                `json:"caption,omitempty"`
	CaptionEntities []*types.MessageEntity `json:"caption_entities,omitempty"`
}

func (c *Captioner) WithCaptionEntities(ce []*types.MessageEntity) {
	c.CaptionEntities = ce
}

func (c *Captioner) WithCaption(caption string) {
	c.Caption = &caption
}
