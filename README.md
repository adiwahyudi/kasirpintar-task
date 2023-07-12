# Kasir Pintar take home test.


## Running 

1. Clone this repository.
2. Install modules `go get`
4. Start server by running `go run main.go`

## Usage
### ***Registration endpoint***
`POST http://127.0.0.1:8080/registration`

**REQUEST**
```json
{
    "iMid": "IONPAYTEST",
    "payMethod": "01",
    "currency": "IDR",
    "amt": "14500",
    "referenceNo": "Adi_001",
    "goodsNm": "Test transaction",
    "dbProcessUrl": "https://merchant.com/api/dbProcessUrl/Notif",
    "userIP": "127.0.0.1",
    "cartData": "{\"count\":1,\"item\":[{\"img_url\":\"https://images.ctfassets.net/od02wyo8cgm5/14Njym0dRLAHaVJywF8WFL/1910357dd0da0ae38b61bc8ad3db86e4/cloudflyer_2-fw19-grey_lime-m-g1.png\",\"goods_name\":\"Shoe\",\"goods_detail\":\"Shoe\",\"goods_amt\":529500}]}",
    "instmntType": 2,
    "instmntMon": 1,
    "recurrOpt": 0,
    "vacctValidDt": "20241212",
    "vacctValidTm": "120000",
    "merFixAcctId": "4",
    "billingNm": "John Doe",
    "billingPhone": "081234567890",
    "billingEmail": "buyer@email.com",
    "billingAddr": "Jl. Perumnas",
    "billingCity": "Bandung",
    "billingState": "Jawa Barat",
    "billingPostCd": "12345",
    "billingCountry": "Indonesia"
}
```
**RESPONSE**
```json
{
    "resultCd": "0000",
    "resultMsg": "SUCCESS",
    "tXid": "IONPAYTEST01202307121613091246",
    "referenceNo": "Adi_001",
    "amt": "14500"
}
```

### ***Status Inquiry endpoint***
`POST http://127.0.0.1:8080/inquiry`

**REQUEST**
```json
{
    "iMid": "IONPAYTEST",
    "tXid": "IONPAYTEST01202307121613091246",
    "referenceNo": "Adi_001",
    "amt": "14500"
}
```
**RESPONSE**
```json
{
    "resultCd": "0000",
    "resultMsg": "init",
    "status": "9",
    "tXid": "IONPAYTEST01202307121613091246",
    "referenceNo": "Adi_001",
    "amt": "14500"
}
```
### ***Payment endpoint***
`POST http://127.0.0.1:8080/payment`

**REQUEST**
```json
{
    "tXid": "IONPAYTEST01202307121613091246",
    "cardNo":"1234567890123450",
    "cardExpYymm":"2412",
    "cardCvv":"141",
    "callBackUrl":"http://localhost:8080/callback",
    "iMid":"IONPAYTEST",
    "referenceNo": "Adi_001",
    "amt": "14500"
}
```
**RESPONSE**
```
    success melakukan payment, namun belum bisa mengeluarkan response dari dalam form pada callBackUrl.
```

Log dapat dilihat pada `log.txt`