package rest

import (
	"encoding/json"
	"goexpert-list-orders/internal/usecase"
	"log"
	"net/http"
)

type Handler struct {
	ListOrdersUC *usecase.ListOrdersUseCase
}

func NewHandler(listOrdersUC *usecase.ListOrdersUseCase) *Handler {
	return &Handler{ListOrdersUC: listOrdersUC}
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ListOrdersUC.Execute(r.Context())
	if err != nil {
		http.Error(w, "Erro ao listar pedidos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		log.Printf("Erro ao codificar resposta: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
	}
}
