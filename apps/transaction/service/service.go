package service

import (
	"context"
	"fmt"

	"github.com/ariefsn/intrans/entities"
	"github.com/ariefsn/intrans/logger"
)

type transactionService struct {
	transactionRepo entities.TransactionRepository
}

// Create implements entities.TransactionService.
func (a *transactionService) Create(ctx context.Context, input entities.TransactionCreatePayload) (*entities.TransactionModel, error) {
	res, err := a.transactionRepo.Create(ctx, input)
	if err != nil {
		logger.Error(fmt.Errorf("create transaction failed. error: %s", err.Error()))
		return nil, err
	}

	logger.Info("create transaction success", entities.M{
		"id": res.ID,
	})

	return res, nil
}

// GetByID implements entities.TransactionService.
func (a *transactionService) GetByID(ctx context.Context, id string) (*entities.TransactionModel, error) {
	res, err := a.transactionRepo.GetByID(ctx, id)
	if err != nil {
		logger.Error(fmt.Errorf("get transaction by id failed. error: %s", err.Error()))
		return nil, err
	}

	logger.Info("get transaction by id success", entities.M{
		"id": res.ID,
	})

	return res, nil

}

func New(transactionRepo entities.TransactionRepository) entities.TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}
