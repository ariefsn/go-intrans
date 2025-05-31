package delivery_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ariefsn/intrans/apps/transaction/delivery"
	ent "github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/helper"
	"github.com/ariefsn/intrans/mocks/entities"
	"github.com/ariefsn/intrans/validator"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(t *testing.T) (svc *entities.TransactionServiceMock, handler *chi.Mux) {
	validator.InitValidator()
	svc = entities.NewTransactionServiceMock(t)
	handler = delivery.New(svc)

	return svc, handler
}

func TestGetByID(t *testing.T) {
	t.Run("Failed getting transaction", func(t *testing.T) {
		svc, r := setup(t)

		id := "1"
		svc.On("GetByID", mock.Anything, id).Return(nil, errors.New("not found"))

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var res ent.TransactionResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Get transaction failed", res.Message)
	})

	t.Run("Success", func(t *testing.T) {
		svc, r := setup(t)

		id := "1"
		mockRes := &ent.TransactionModel{
			ID:                   uuid.New(),
			SourceAccountID:      1,
			DestinationAccountID: 2,
			Amount:               100,
			CreatedAt:            time.Now(),
		}
		svc.On("GetByID", mock.Anything, id).Return(mockRes, nil)

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var res ent.TransactionResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.True(t, res.Status)
		assert.Equal(t, mockRes.ID, res.Data.ID)
		assert.Equal(t, mockRes.SourceAccountID, res.Data.SourceAccountID)
		assert.Equal(t, mockRes.DestinationAccountID, res.Data.DestinationAccountID)
		assert.Equal(t, mockRes.Amount, res.Data.Amount)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Failed wrong payload", func(t *testing.T) {
		_, r := setup(t)

		payload, _ := helper.ToJsonBody(ent.M{
			"random_field": "random_value",
		})

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var res ent.TransactionResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Invalid payload", res.Message)
	})

	t.Run("Failed missing mandatory field", func(t *testing.T) {
		_, r := setup(t)

		payload, _ := helper.ToJsonBody(ent.M{
			"amount": 100,
		})

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var res ent.TransactionResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Invalid payload", res.Message)
	})

	t.Run("Failed getting transaction", func(t *testing.T) {
		svc, r := setup(t)

		payloadStruct := ent.TransactionCreatePayload{
			SourceAccountID:      1,
			DestinationAccountID: 2,
			Amount:               100,
		}
		payload, _ := helper.ToJsonBody(payloadStruct)

		svc.On("Create", mock.Anything, payloadStruct).Return(nil, errors.New("not found"))

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var res ent.TransactionResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Create transaction failed", res.Message)
	})

	t.Run("Success", func(t *testing.T) {
		svc, r := setup(t)

		payloadStruct := ent.TransactionCreatePayload{
			SourceAccountID:      1,
			DestinationAccountID: 2,
			Amount:               100,
		}
		payload, _ := helper.ToJsonBody(payloadStruct)

		mockRes := &ent.TransactionModel{
			ID:                   uuid.New(),
			SourceAccountID:      1,
			DestinationAccountID: 2,
			Amount:               100,
			CreatedAt:            time.Now(),
		}
		svc.On("Create", mock.Anything, payloadStruct).Return(mockRes, nil)

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var res ent.TransactionResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.True(t, res.Status)
		assert.Equal(t, mockRes.ID, res.Data.ID)
		assert.Equal(t, mockRes.SourceAccountID, res.Data.SourceAccountID)
		assert.Equal(t, mockRes.DestinationAccountID, res.Data.DestinationAccountID)
		assert.Equal(t, mockRes.Amount, res.Data.Amount)
	})
}
