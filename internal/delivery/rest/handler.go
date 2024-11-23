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

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Obtém o contexto da requisição
	orders, err := h.ListOrdersUC.Execute(ctx)
	if err != nil {
		log.Printf("Erro ao listar pedidos: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
