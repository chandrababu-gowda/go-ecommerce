package products

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/chandrababu-gowda/go-ecommerce/internal/json"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{service: s}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("Failed in handler", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)

}

func (h *handler) FindProductByID(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	product, err := h.service.FindProductByID(r.Context(), id)
	if err != nil {
		slog.Error("Failed to fetch product by id", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
