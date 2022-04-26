package types

type ProtectContenter struct {
	ProtectContent bool `json:"protect_content,omitempty"`
}

func (pc *ProtectContenter) WithProtectContent() {
	pc.ProtectContent = true
}
