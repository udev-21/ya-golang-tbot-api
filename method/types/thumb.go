package types

type Thumber struct {
	Thumb InputFile `json:"thumb,omitempty"`
}

func (t *Thumber) WithThumb(thumb InputFile) {
	t.Thumb = thumb
}
