package handler

import (
    "encoding/json"
    "net/http"
    "go_race_condition/src/usecases"
)

type Handler struct {
    productUsecase usecase.ProductUsecase
}

func NewHandler(productUsecase usecase.ProductUsecase) *Handler {
    return &Handler{productUsecase: productUsecase}
}

func (h *Handler) GetStockHandler(w http.ResponseWriter, r *http.Request) {
    stock := h.productUsecase.GetStock()
    response := struct {
        Stock int `json:"stock"`
    }{
        Stock: stock,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateStockWithWaitGroupHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Amount int `json:"amount"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    h.productUsecase.UpdateStockWithWaitGroup(req.Amount)
    h.GetStockHandler(w, r)
}

func (h *Handler) UpdateStockWithMutexHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Amount int `json:"amount"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    h.productUsecase.UpdateStockWithMutex(req.Amount)
    h.GetStockHandler(w, r)
}
