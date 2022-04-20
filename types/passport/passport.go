package passport

type Passport struct {
	MustHaveIdDocumentData
	MustHaveFrontSide

	FrontSide File `json:"front_side"`

	MightHaveSelfie
	MightHaveTranslation
}
