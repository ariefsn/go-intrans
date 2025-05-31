package delivery

import (
	"net/http"
	"strconv"

	"github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/helper"
	"github.com/ariefsn/intrans/logger"
	"github.com/ariefsn/intrans/validator"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service entities.AccountService
}

// @Tags Account
// @Summary Get account by ID
// @Description Get account by ID
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} entities.AccountResponse
// @Router /accounts/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		logger.Error(err)
		res := entities.Response{
			Status:  false,
			Message: "Invalid ID",
			Errors:  []string{"id must be a number"},
		}
		res.Send(w, http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetByID(r.Context(), int64(idNum))
	if err != nil {
		res := entities.Response{
			Status:  false,
			Message: "Get account failed",
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

// @Tags Account
// @Summary Create a new account
// @Description Create a new account
// @Accept json
// @Produce json
// @Param payload body entities.AccountCreatePayload true "Payload"
// @Success 200 {object} entities.AccountResponse
// @Router /accounts [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload entities.AccountCreatePayload
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
			Message: "Create account failed",
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

func New(service entities.AccountService) *chi.Mux {
	h := &Handler{
		Service: service,
	}

	r := chi.NewRouter()

	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)

	return r
}
