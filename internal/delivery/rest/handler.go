package rest

import (
	"encoding/json"
	"goexpert-list-orders/internal/domain"
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

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	id, err := h.ListOrdersUC.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, "Erro ao criar pedido", http.StatusInternalServerError)
		return
	}

	order.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
