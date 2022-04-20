package passport

type IDDocumentData struct {
	DocumentNo string  `json:"document_no"`
	ExpiryDate *string `json:"expiry_date"`
}
