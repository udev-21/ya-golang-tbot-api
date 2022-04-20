package passport

type SecureValue struct {
	Data        *DataCredentials  `json:"data,omitempty"`
	FrontSide   *FileCredentials  `json:"front_side,omitempty"`
	ReverseSide *FileCredentials  `json:"reverse_side,omitempty"`
	Selfie      *FileCredentials  `json:"selfie,omitempty"`
	Translation []FileCredentials `json:"translation,omitempty"`
	Files       []FileCredentials `json:"files,omitempty"`
}
