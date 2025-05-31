package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ariefsn/intrans/apps/account/service"
	ent "github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/mocks/entities"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (repo *entities.AccountRepositoryMock, svc ent.AccountService) {
	repo = entities.NewAccountRepositoryMock(t)
	svc = service.New(repo)

	return repo, svc
}

func TestGetByID(t *testing.T) {
	t.Run("Failed", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		id := int64(1)
		repo.On("GetByID", ctx, id).Return(nil, errors.New("not found"))
		res, err := svc.GetByID(ctx, id)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("Success", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		id := int64(1)
		mockRes := &ent.AccountModel{ID: id, Balance: 1000}
		repo.On("GetByID", ctx, id).Return(mockRes, nil)
		res, err := svc.GetByID(ctx, id)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockRes)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Failed", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		input := ent.AccountCreatePayload{ID: 2, InitialBalance: 123}
		repo.On("Create", ctx, input).Return(nil, errors.New("not found"))
		res, err := svc.Create(ctx, input)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("Success", func(t *testing.T) {
		repo, svc := setup(t)

		ctx := context.Background()
		input := ent.AccountCreatePayload{ID: 2, InitialBalance: 123}
		mockRes := &ent.AccountModel{ID: input.ID, Balance: input.InitialBalance}
		repo.On("Create", ctx, input).Return(mockRes, nil)
		res, err := svc.Create(ctx, input)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockRes)
	})
}
