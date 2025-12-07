package rest

import (
	"encoding/json"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/model"
)

type RefundPaymentRequest struct {
	TransactionID int64 `json:"transaction_id"`
}

func (s *Server) RefundPayment(w http.ResponseWriter, r *http.Request) {
	reqFromBody := &RefundPaymentRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqFromBody); err != nil {
		MakeErrorResponse(w, err, http.StatusBadRequest)
	}

	req := model.RefundPaymentRequest{
		TransactionID: reqFromBody.TransactionID,
	}

	err := s.refundPaymentUsecase.Execute(r.Context(), req.TransactionID)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusTeapot)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
