package passport

type InternalPassport struct {
	MustHaveIdDocumentData
	MustHaveFrontSide

	MightHaveSelfie
	MightHaveTranslation
}
