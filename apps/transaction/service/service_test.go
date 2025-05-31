package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ariefsn/intrans/apps/transaction/service"
	ent "github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/mocks/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (repo *entities.TransactionRepositoryMock, svc ent.TransactionService) {
	repo = entities.NewTransactionRepositoryMock(t)
	svc = service.New(repo)

	return repo, svc
}

func TestGetByID(t *testing.T) {
	t.Run("Failed", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		id := uuid.New()
		repo.On("GetByID", ctx, id.String()).Return(nil, errors.New("not found"))
		res, err := svc.GetByID(ctx, id.String())

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("Success", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		id := uuid.New()
		mockRes := &ent.TransactionModel{ID: id, SourceAccountID: 1, DestinationAccountID: 2, Amount: 1000}
		repo.On("GetByID", ctx, id.String()).Return(mockRes, nil)
		res, err := svc.GetByID(ctx, id.String())

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockRes)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Failed", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		input := ent.TransactionCreatePayload{SourceAccountID: 1, DestinationAccountID: 2, Amount: 1000}
		repo.On("Create", ctx, input).Return(nil, errors.New("not found"))
		res, err := svc.Create(ctx, input)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("Success", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		input := ent.TransactionCreatePayload{SourceAccountID: 1, DestinationAccountID: 2, Amount: 1000}
		mockRes := &ent.TransactionModel{ID: uuid.New(), SourceAccountID: input.SourceAccountID, DestinationAccountID: input.DestinationAccountID, Amount: input.Amount}
		repo.On("Create", ctx, input).Return(mockRes, nil)
		res, err := svc.Create(ctx, input)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockRes)
	})
}
