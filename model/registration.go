package model

type RegistrationRequest struct {
	Timestamp          string `json:"timeStamp"`
	MerchantId         string `json:"iMid"`
	PaymentMethod      string `json:"payMethod"`
	Currency           string `json:"currency"`
	Amount             string `json:"amt"`
	ReferenceNo        string `json:"referenceNo"`
	GoodsName          string `json:"goodsNm"`
	BuyerName          string `json:"billingNm"`
	BuyerPhone         string `json:"billingPhone"`
	BuyerEmail         string `json:"billingEmail"`
	BuyerAddress       string `json:"billingAddr"`
	BuyerCity          string `json:"billingCity"`
	BillingState       string `json:"billingState"`
	BillingPostNumber  string `json:"billingPostCd"`
	BillingCountry     string `json:"billingCountry"`
	CartData           string `json:"cartData"`
	InstallmentType    int    `json:"instmntType"`
	InstallmentMonth   int    `json:"instmntMon"`
	RecurringOption    int    `json:"recurrOpt"`
	VaValidDate        string `json:"vacctValidDt"`
	VaValidTime        string `json:"vacctValidTm"`
	MerchantReservedId string `json:"merFixAcctId"`
	UserIp             string `json:"userIP"`
	NotificationUrl    string `json:"dbProcessUrl"`
	MerchantToken      string `json:"merchantToken"`
}

type RegistrationResponse struct {
	Code          string `json:"resultCd"`
	Message       string `json:"resultMsg"`
	TransactionId string `json:"tXid"`
	ReferenceNo   string `json:"referenceNo"`
	Amount        string `json:"amt"`
}
