package delivery

import (
	"net/http"

	"github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/helper"
	"github.com/ariefsn/intrans/logger"
	"github.com/ariefsn/intrans/validator"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service entities.TransactionService
}

// @Tags Transaction
// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} entities.TransactionResponse
// @Router /transactions/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		res := entities.Response{
			Status:  false,
			Message: "Get transaction failed",
			Errors:  []string{err.Error()},
		}
		res.Send(w, http.StatusInternalServerError)
		return
	}

	res := entities.Response{
		Status: true,
		Data:   data,
	}
	res.Send(w, http.StatusOK)
}

// @Tags Transaction
// @Summary Create a new transaction
// @Description Create a new transaction
// @Accept json
// @Produce json
// @Param payload body entities.TransactionCreatePayload true "Payload"
// @Success 200 {object} entities.TransactionResponse
// @Router /transactions [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload entities.TransactionCreatePayload
	if err := helper.ParsePayload(r, &payload); err != nil {
		logger.Error(err)
		res := entities.Response{
			Status:  false,
			Message: "Invalid payload",
			Errors:  []string{err.Error()},
		}
		res.Send(w, http.StatusBadRequest)
		return
	}

	if err, msg := validator.ValidateStruct(payload); err != nil {
		logger.Error(err)
		res := entities.Response{
			Status:  false,
			Message: "Invalid payload",
			Errors:  msg,
		}
		res.Send(w, http.StatusBadRequest)
		return
	}

	data, err := h.Service.Create(r.Context(), payload)
	if err != nil {
		res := entities.Response{
			Status:  false,
			Message: "Create transaction failed",
			Errors:  []string{err.Error()},
		}
		res.Send(w, http.StatusInternalServerError)
		return
	}

	res := entities.Response{
		Status: true,
		Data:   data,
	}
	res.Send(w, http.StatusOK)
}

func New(service entities.TransactionService) *chi.Mux {
	h := &Handler{
		Service: service,
	}

	r := chi.NewRouter()

	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)

	return r
}
