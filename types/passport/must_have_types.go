package passport

type MustHaveReverseSide struct {
	ReverseSide File `json:"reverse_side"`
}

type MustHaveFrontSide struct {
	FrontSide File `json:"front_side"`
}

type MustHaveFiles struct {
	Files []File `json:"files"`
}

// according to https://core.telegram.org/passport#fields
// key of all IDDocumentData types are `data`
// except types: personal_details and address
type MustHaveIdDocumentData struct {
	Data IDDocumentData `json:"data"`
}
