package passport

type MightHaveTranslation struct {
	Translation []File `json:"translation,omitempty"`
}

type MightHaveSelfie struct {
	Selfie *File `json:"selfie,omitempty"`
}
