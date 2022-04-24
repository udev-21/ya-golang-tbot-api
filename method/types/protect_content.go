package types

type ProtectContenter struct {
	ProtectContent_ bool `json:"protect_content,omitempty"`
}

func (pc *ProtectContenter) WithProtectContent() {
	pc.ProtectContent_ = true
}
