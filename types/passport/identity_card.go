package passport

type IdentityCard struct {
	MustHaveIdDocumentData
	MustHaveFrontSide
	MustHaveReverseSide

	MightHaveSelfie
	MightHaveTranslation
}
