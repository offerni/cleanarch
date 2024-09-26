package web

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func (h *WebOrderHandler) Fetch(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing order ID", http.StatusBadRequest)
		return
	}

	fetchOrder := usecase.NewFetchOrderUseCase(h.OrderRepository)
	output, err := fetchOrder.Execute(usecase.FetchOrderInputDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
