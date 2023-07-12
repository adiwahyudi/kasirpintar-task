package model

type PaymentRequest struct {
	TransactionId    string `json:"tXid"`
	CcNumber         string `json:"cardNo"`
	CcExpiry         string `json:"cardExpYymm"`
	Cvv              string `json:"cardCvv"`
	PaymentResultUrl string `json:"callBackUrl"`
	MerchantId       string `json:"iMid"`
	Amount           string `json:"amt"`
	ReferenceNo      string `json:"referenceNo"`
}

type PaymentSent struct {
	Timestamp        string `json:"timeStamp"`
	TransactionId    string `json:"tXid"`
	CcNumber         string `json:"cardNo"`
	CcExpiry         string `json:"cardExpYymm"`
	Cvv              string `json:"cardCvv"`
	PaymentResultUrl string `json:"callBackUrl"`
	MerchantToken    string `json:"merchantToken"`
}
