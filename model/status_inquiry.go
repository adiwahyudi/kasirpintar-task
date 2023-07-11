package model

type StatusInquiryRequest struct {
	Timestamp     string `json:"timeStamp"`
	TransactionId string `json:"tXid"`
	MerchantId    string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amount        string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

type StatusInquiryResponse struct {
	Code          string `json:"resultCd"`
	Message       string `json:"resultMsg"`
	Status        string `json:"status"`
	TransactionId string `json:"tXid"`
	ReferenceNo   string `json:"referenceNo"`
	Amount        string `json:"amt"`
}
