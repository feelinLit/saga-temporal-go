package rest

import (
	"encoding/json"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/model"
)

type AuthorizePaymentRequest struct {
	AccountID int64 `json:"account_id"`
	Amount    int32 `json:"amount"`
}

type AuthorizePaymentResponse struct {
	TransactionID int64 `json:"transaction_id"`
}

func (s *Server) AuthorizePayment(w http.ResponseWriter, r *http.Request) {
	reqFromBody := &AuthorizePaymentRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqFromBody); err != nil {
		MakeErrorResponse(w, err, http.StatusBadRequest)
	}

	req := model.AuthorizePaymentRequest{
		AccountID: reqFromBody.AccountID,
		Amount:    reqFromBody.Amount,
	}

	transactionID, err := s.authorizePaymentUsecase.Execute(r.Context(), req)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusTeapot)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &AuthorizePaymentResponse{TransactionID: transactionID}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		MakeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
}
