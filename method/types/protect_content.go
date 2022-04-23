package types

type ProtectContent struct {
	ProtectContent_ bool `json:"protect_content,omitempty"`
}

func (pc *ProtectContent) ProtectContent() {
	pc.ProtectContent_ = true
}
