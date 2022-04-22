package types

import "github.com/udev21/golang-tbot-api/utils"

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	Selective             *bool   `json:"selective,omitempty"`
}

func (fr *ForceReply) RawJsonPayload() (map[string]interface{}, error) {
	return utils.ConvertToMapStringInterface(fr)
}
func (fr *ForceReply) WithSelective() {
	fr.Selective = new(bool)
	*fr.Selective = true
}
func (fr *ForceReply) WithInputFieldPlaceholder(placeholder string) {
	fr.InputFieldPlaceholder = new(string)
	fr.InputFieldPlaceholder = &placeholder
}

func NewForceReply() *ForceReply {
	return &ForceReply{
		ForceReply: true,
	}
}
