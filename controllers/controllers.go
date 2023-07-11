package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kasirpintar/helper"
	"kasirpintar/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Registration(ctx *gin.Context) {
	url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"

	newRegistration := model.RegistrationRequest{}

	// Binding to JSON
	if err := ctx.ShouldBindJSON(&newRegistration); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	t := time.Now()
	newRegistration.Timestamp = t.Format("20060102150405")
	newRegistration.MerchantId = os.Getenv("MERCHANT_ID")

	// SHA256 Merchant Token
	merchantKey := os.Getenv("MERCHANT_KEY")
	merchantToken := newRegistration.Timestamp + newRegistration.MerchantId + newRegistration.ReferenceNo + newRegistration.Amount + merchantKey
	merchantToken = helper.SHA256toString(merchantToken)
	newRegistration.MerchantToken = merchantToken
	request, err := json.Marshal(newRegistration)

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

	if err != nil {
		fmt.Println(err)
	}

	response := model.RegistrationResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
	}

	// Return response
	ctx.JSON(result.StatusCode, response)
}
