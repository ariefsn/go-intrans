package service

import (
	"context"
	"fmt"

	"github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/logger"
)

type accountService struct {
	accountRepo entities.AccountRepository
}

// Create implements entities.AccountService.
func (a *accountService) Create(ctx context.Context, input entities.AccountCreatePayload) (*entities.AccountModel, error) {
	res, err := a.accountRepo.Create(ctx, input)
	if err != nil {
		logger.Error(fmt.Errorf("create account failed. error: %s", err.Error()))
		return nil, err
	}

	logger.Info("create account success", entities.M{
		"id": res.ID,
	})

	return res, nil
}

// GetByID implements entities.AccountService.
func (a *accountService) GetByID(ctx context.Context, id int64) (*entities.AccountModel, error) {
	res, err := a.accountRepo.GetByID(ctx, id)
	if err != nil {
		logger.Error(fmt.Errorf("get account by id failed. error: %s", err.Error()))
		return nil, err
	}

	logger.Info("get account by id success", entities.M{
		"id": res.ID,
	})

	return res, nil
}

func New(accountRepo entities.AccountRepository) entities.AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}
