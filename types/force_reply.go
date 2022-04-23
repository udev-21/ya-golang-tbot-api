package types

func NewForceReply() *ForceReply {
	return &ForceReply{
		ForceReply: true,
	}
}

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	Selective             *bool   `json:"selective,omitempty"`
}

func (fr *ForceReply) WithSelective() {
	fr.Selective = new(bool)
	*fr.Selective = true
}
func (fr *ForceReply) WithInputFieldPlaceholder(placeholder string) {
	fr.InputFieldPlaceholder = new(string)
	fr.InputFieldPlaceholder = &placeholder
}

// IsReplyMarkup - to satisfy ReplyMarkup interface
func (fr *ForceReply) IsReplyMarkup() {}
