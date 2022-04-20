package passport

type DriverLicense struct {
	MustHaveIdDocumentData
	MustHaveFrontSide
	MustHaveReverseSide
	MightHaveSelfie
	MightHaveTranslation
}
