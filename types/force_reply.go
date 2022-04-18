package types

type ForceReply struct {
	ForceReply            bool  `json:"force_reply"`
	InputFieldPlaceholder *bool `json:"input_field_placeholder,omitempty"`
	Selective             *bool `json:"selective,omitempty"`
}
