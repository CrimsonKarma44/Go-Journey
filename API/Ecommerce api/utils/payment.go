package utils

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"os"
)

import (
	"bytes"
	"io/ioutil"
)

type PaymentHandler struct {
	DB *gorm.DB
}

func (p *PaymentHandler) InitializePayment(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var requestBody struct {
		Email  string `json:"email"`
		Amount int    `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Prepare the request payload
	payload := map[string]interface{}{
		"email":  requestBody.Email,
		"amount": requestBody.Amount,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to marshal payload", http.StatusInternalServerError)
		return
	}

	// Make a POST request to Paystack's API
	req, err := http.NewRequest("POST", paystackBaseURL+"/transaction/initialize", bytes.NewBuffer(payloadBytes))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+paystackSecretKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make request to Paystack", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Parse the response
	var paystackResp InitializeTransactionResponse
	err = json.Unmarshal(body, &paystackResp)
	if err != nil {
		http.Error(w, "Failed to parse Paystack response", http.StatusInternalServerError)
		return
	}

	// Check if the transaction was initialized successfully
	if !paystackResp.Status {
		http.Error(w, paystackResp.Message, http.StatusBadRequest)
		return
	}

	// Return the authorization URL to the client
	response := map[string]string{
		"authorization_url": paystackResp.Data.AuthorizationURL,
		"reference":         paystackResp.Data.Reference,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

const paystackBaseURL = "https://api.paystack.co"

var paystackSecretKey = os.Getenv("PAYSTACK_SECRET_KEY")

type InitializeTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

type VerifyTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Status    string `json:"status"`
		Reference string `json:"reference"`
		Amount    int    `json:"amount"`
		PaidAt    string `json:"paid_at"`
		Channel   string `json:"channel"`
		Currency  string `json:"currency"`
	} `json:"data"`
}
