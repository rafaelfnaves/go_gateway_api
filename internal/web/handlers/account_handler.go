package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelfnaves/go-gateway-api/internal/dto"
	"github.com/rafaelfnaves/go-gateway-api/internal/service"
)

type AccountHandler struct {
	accountservice service.AccountService
}

func NewAccountHandler(accountservice service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountservice: accountservice,
	}
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.accountservice.CreateAccount(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
