package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const paymentServiceAddress = "http://localhost:8072"

type PaymentService struct {
	httpClient *http.Client
}

func NewPaymentService(httpClient *http.Client) *PaymentService {
	return &PaymentService{httpClient: httpClient}
}

func (s *PaymentService) AuthorizePayment(accountID int64, amount int32) (transactionID int64, error error) {
	data := struct {
		AccountID int64 `json:"account_id"`
		Amount    int32 `json:"amount"`
	}{
		AccountID: accountID,
		Amount:    amount,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return -1, fmt.Errorf("json.Marshal: %w", err)
	}

	url := fmt.Sprintf("%s/%s", paymentServiceAddress, "payment")
	contentType := "application/json"
	resp, err := s.httpClient.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return -1, fmt.Errorf("httpClient.Post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("expected status code: 200; received: %d", resp.StatusCode)
	}

	var authorizedPayment struct {
		TransactionID int64 `json:"transaction_id"`
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("failed reading response body: %w", err)
	}
	if err = json.Unmarshal(respBodyBytes, &authorizedPayment); err != nil {
		return -1, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return authorizedPayment.TransactionID, nil
}

func (s *PaymentService) RefundPayment(transactionID int64) error {
	data := struct {
		TransactionID int64 `json:"transaction_id"`
	}{
		TransactionID: transactionID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	url := fmt.Sprintf("%s/%s", paymentServiceAddress, "payment/refund")
	contentType := "application/json"
	resp, err := s.httpClient.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("httpClient.Post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code: 200; received: %d", resp.StatusCode)
	}

	return nil
}
