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
	"time"

	"github.com/gin-gonic/gin"
)

func Registration(ctx *gin.Context) {
	log.Println("Registration")
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

	log.Println("Request : ", newRequest)
	request, err := json.Marshal(newRequest)
	if err != nil {
		fmt.Println(err)
	}
	requestAsBytes := bytes.NewBuffer(request)

	// Request to URL
	result, err := http.Post(url, "application/json", requestAsBytes)

	if err != nil {
		fmt.Println(err)
	}

	defer result.Body.Close()

	// Getting Response
	body, err := io.ReadAll(result.Body)

	if err != nil {
		fmt.Println(err)
	}
	log.Println("Response : ", string(body))
	response := model.RegistrationResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
	}

	// Return response
	ctx.JSON(result.StatusCode, response)
}

func StatusInquiry(ctx *gin.Context) {
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
	fmt.Println(newRequest)
	if err != nil {
		fmt.Println(err)
	}
	requestAsBytes := bytes.NewBuffer(request)

	// Request
	result, err := http.Post(url, "application/json", requestAsBytes)

	if err != nil {
		fmt.Println(err)
	}

	defer result.Body.Close()

	// Getting Response
	body, err := io.ReadAll(result.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
	}

	response := model.StatusInquiryResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
	}

	// Return response
	ctx.JSON(result.StatusCode, response)
}

func Callback(ctx *gin.Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error parsing form data")
		return
	}

	for key, values := range ctx.Request.PostForm {
		for _, value := range values {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}
	}

	// Send a response
	ctx.String(http.StatusOK, "Callback received")
	fmt.Println("Ini Callback")
}

func Payment(ctx *gin.Context) {

	endpoint := "https://dev.nicepay.co.id/nicepay/direct/v2/payment?"

	newPayment := model.PaymentRequest{}
	if err := ctx.Bind(&newPayment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	timeStamp := time.Now().Format("20060102150405")
	merchantKey := os.Getenv("MERCHANT_KEY")
	merchantToken := timeStamp + newPayment.MerchantId + newPayment.ReferenceNo + newPayment.Amount + merchantKey
	merchantToken = helper.SHA256toString(merchantToken)

	sendToNicePay := model.PaymentSent{
		Timestamp:        timeStamp,
		TransactionId:    newPayment.TransactionId,
		CcNumber:         newPayment.CcNumber,
		CcExpiry:         newPayment.CcExpiry,
		Cvv:              newPayment.Cvv,
		PaymentResultUrl: newPayment.PaymentResultUrl,
		MerchantToken:    merchantToken,
	}

	requestJson, _ := json.Marshal(sendToNicePay)
	requestMap := map[string]string{}
	json.Unmarshal(requestJson, &requestMap)
	data := url.Values{}
	for key, val := range requestMap {
		data.Add(key, val)
	}
	queryParams := data.Encode()
	newUrl := endpoint + queryParams

	ctx.Redirect(http.StatusFound, newUrl)
}
