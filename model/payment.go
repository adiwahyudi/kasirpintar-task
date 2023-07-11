package model

type PaymentRequest struct {
	TransactionId    string `form:"tXid"`
	CcNumber         string `form:"cardNo"`
	CcExpiry         string `form:"cardExpYymm"`
	Cvv              string `form:"cardCvv"`
	PaymentResultUrl string `form:"callBackUrl"`
	MerchantId       string `form:"iMid"`
	Amount           string `form:"amt"`
	ReferenceNo      string `form:"referenceNo"`
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

type PaymentRespone struct {
	Code             string `json:"resultCd"`
	Message          string `json:"resultMsg"`
	TransactionId    string `json:"tXid"`
	ReferenceNo      string `json:"referenceNo"`
	PaymentMethod    string `json:"payMethod"`
	Currency         string `json:"currency"`
	Amount           string `json:"amt"`
	GoodsName        string `json:"goodsNm"`
	BuyerName        string `json:"billingNm"`
	TransactionDate  string `json:"transDt"`
	TransactionTime  string `json:"transTm"`
	Description      string `json:"description"`
	ApprovalNumber   string `json:"authNo"`
	IssuerBankCode   string `json:"issuBankCd"`
	AcquireBankCode  string `json:"acquBankCd"`
	MaskedCard       string `json:"cardNo"`
	CardExpiry       string `json:"cardExpYymm"`
	InstallmentType  string `json:"instmntType"`
	InstallmentMonth string `json:"instmntMon"`
	RecurringOption  string `json:"recurrOpt"`
	CcTrans          string `json:"ccTransType"`
}
