package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kasirpintar/helper"
	"kasirpintar/model"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Registration(ctx *gin.Context) {
	log.Println("Doing registration...")
	url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"

	newRequest := model.RegistrationRequest{}
	// Binding to JSON
	if err := ctx.ShouldBindJSON(&newRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	newRequest.Timestamp = time.Now().Format("20060102150405")

	// SET MERCHANT TOKEN with SHA256
	merchantKey := os.Getenv("MERCHANT_KEY")
	merchantToken := newRequest.Timestamp + newRequest.MerchantId + newRequest.ReferenceNo + newRequest.Amount + merchantKey
	merchantToken = helper.SHA256toString(merchantToken)
	newRequest.MerchantToken = merchantToken

	request, err := json.Marshal(newRequest)
	log.Println("Registration request : ", string(request))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	requestAsBytes := bytes.NewBuffer(request)
	// Request to URL
	result, err := http.Post(url, "application/json", requestAsBytes)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer result.Body.Close()

	// Getting Response
	body, err := io.ReadAll(result.Body)

	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("Registration response : ", string(body))
	response := model.RegistrationResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Return response
	ctx.JSON(result.StatusCode, response)
}

func StatusInquiry(ctx *gin.Context) {
	log.Println("Check inquiry status...")
	url := "https://dev.nicepay.co.id/nicepay/direct/v2/inquiry"

	newRequest := model.StatusInquiryRequest{}
	// Binding to JSON
	if err := ctx.ShouldBindJSON(&newRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	t := time.Now()
	newRequest.Timestamp = t.Format("20060102150405")

	// SHA256 Merchant Token
	merchantKey := os.Getenv("MERCHANT_KEY")
	merchantToken := newRequest.Timestamp + newRequest.MerchantId + newRequest.ReferenceNo + newRequest.Amount + merchantKey
	merchantToken = helper.SHA256toString(merchantToken)
	newRequest.MerchantToken = merchantToken

	request, err := json.Marshal(newRequest)
	log.Println("Status inquiry request : ", string(request))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	requestAsBytes := bytes.NewBuffer(request)

	// Request
	result, err := http.Post(url, "application/json", requestAsBytes)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer result.Body.Close()

	// Getting Response
	body, err := io.ReadAll(result.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("Status inquiry response : ", string(body))
	response := model.StatusInquiryResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Return response
	ctx.JSON(result.StatusCode, response)
}

func Payment(ctx *gin.Context) {
	log.Println("Doing payment...")

	endpoint := "https://dev.nicepay.co.id/nicepay/direct/v2/payment?"

	newPayment := model.PaymentRequest{}
	if err := ctx.ShouldBindJSON(&newPayment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	timeStamp := time.Now().Format("20060102150405")
	merchantKey := os.Getenv("MERCHANT_KEY")
	merchantToken := timeStamp + newPayment.MerchantId + newPayment.ReferenceNo + newPayment.Amount + merchantKey
	merchantToken = helper.SHA256toString(merchantToken)

	newPaymentJson, _ := json.Marshal(newPayment)
	log.Println("Payment request => ", string(newPaymentJson))

	sendToNicePay := model.PaymentSent{
		Timestamp:        timeStamp,
		TransactionId:    newPayment.TransactionId,
		CcNumber:         newPayment.CcNumber,
		CcExpiry:         newPayment.CcExpiry,
		Cvv:              newPayment.Cvv,
		PaymentResultUrl: newPayment.PaymentResultUrl,
		MerchantToken:    merchantToken,
	}

	// Struct -> JSON -> Map, to set the encoded form data.
	requestJson, _ := json.Marshal(sendToNicePay)
	requestMap := map[string]string{}
	json.Unmarshal(requestJson, &requestMap)
	data := url.Values{}
	for key, val := range requestMap {
		data.Add(key, val)
	}

	// Hit Nicepay Payment Endpoint
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	respBody, _ := io.ReadAll(resp.Body)

	log.Println("Payment response => \n", string(respBody))

}

// func Callback(ctx *gin.Context) {
// 	err := ctx.Request.ParseForm()
// 	if err != nil {
// 		ctx.String(http.StatusBadRequest, "Error parsing form data")
// 		return
// 	}

// 	for key, values := range ctx.Request.PostForm {
// 		for _, value := range values {
// 			fmt.Printf("Key: %s, Value: %s\n", key, value)
// 		}
// 	}

// 	// Send a response
// 	ctx.String(http.StatusOK, "Callback received")
// 	fmt.Println("Ini Callback")
// }
