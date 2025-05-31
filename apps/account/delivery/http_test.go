package delivery_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariefsn/intrans/apps/account/delivery"
	ent "github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/helper"
	"github.com/ariefsn/intrans/mocks/entities"
	"github.com/ariefsn/intrans/validator"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(t *testing.T) (svc *entities.AccountServiceMock, handler *chi.Mux) {
	validator.InitValidator()
	svc = entities.NewAccountServiceMock(t)
	handler = delivery.New(svc)

	return svc, handler
}

func TestGetByID(t *testing.T) {
	t.Run("Failed id not a number", func(t *testing.T) {
		_, r := setup(t)

		req := httptest.NewRequest(http.MethodGet, "/a", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Invalid ID", res.Message)
	})

	t.Run("Failed getting account", func(t *testing.T) {
		svc, r := setup(t)

		id := int64(1)
		svc.On("GetByID", mock.Anything, id).Return(nil, errors.New("not found"))

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Get account failed", res.Message)
	})

	t.Run("Success", func(t *testing.T) {
		svc, r := setup(t)

		id := int64(1)
		mockRes := &ent.AccountModel{
			ID:      1,
			Balance: 1000,
		}
		svc.On("GetByID", mock.Anything, id).Return(mockRes, nil)

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.True(t, res.Status)
		assert.Equal(t, mockRes.ID, res.Data.ID)
		assert.Equal(t, mockRes.Balance, res.Data.Balance)
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

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Invalid payload", res.Message)
	})

	t.Run("Failed missing mandatory field", func(t *testing.T) {
		_, r := setup(t)

		payload, _ := helper.ToJsonBody(ent.M{
			"initial_balance": 100,
		})

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Invalid payload", res.Message)
	})

	t.Run("Failed getting account", func(t *testing.T) {
		svc, r := setup(t)

		payloadStruct := ent.AccountCreatePayload{
			ID:             1,
			InitialBalance: 100,
		}
		payload, _ := helper.ToJsonBody(payloadStruct)

		svc.On("Create", mock.Anything, payloadStruct).Return(nil, errors.New("not found"))

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.False(t, res.Status)
		assert.Equal(t, "Create account failed", res.Message)
	})

	t.Run("Success", func(t *testing.T) {
		svc, r := setup(t)

		payloadStruct := ent.AccountCreatePayload{
			ID:             1,
			InitialBalance: 100,
		}
		payload, _ := helper.ToJsonBody(payloadStruct)

		mockRes := &ent.AccountModel{
			ID:      1,
			Balance: 1000,
		}
		svc.On("Create", mock.Anything, payloadStruct).Return(mockRes, nil)

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var res ent.AccountResponse
		err := json.Unmarshal(rec.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.True(t, res.Status)
		assert.Equal(t, mockRes.ID, res.Data.ID)
		assert.Equal(t, mockRes.Balance, res.Data.Balance)
	})
}
