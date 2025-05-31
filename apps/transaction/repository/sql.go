package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ariefsn/intrans/entities"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type transactionRepository struct {
	Db *bun.DB
}

// Create implements entities.TransactionRepository.
func (a *transactionRepository) Create(ctx context.Context, input entities.TransactionCreatePayload) (*entities.TransactionModel, error) {
	data := &entities.TransactionModel{
		ID:                   uuid.New(),
		SourceAccountID:      input.SourceAccountID,
		DestinationAccountID: input.DestinationAccountID,
		Amount:               input.Amount,
		CreatedAt:            time.Now(),
	}

	err := a.Db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// Get the sender
		var from entities.AccountModel
		if err := a.Db.NewSelect().Model(&from).Where("id = ?", input.SourceAccountID).Scan(ctx); err != nil {
			return err
		}

		if from.Balance < input.Amount {
			return errors.New("insuffience balance")
		}

		// Update the sender
		if _, err := a.Db.NewUpdate().Model(&entities.AccountModel{}).Set("balance = balance - ?", input.Amount).Where("id = ?", input.SourceAccountID).Exec(ctx); err != nil {
			return err
		}

		// Update the recipient
		if _, err := a.Db.NewUpdate().Model(&entities.AccountModel{}).Set("balance = balance + ?", input.Amount).Where("id = ?", input.DestinationAccountID).Exec(ctx); err != nil {
			return err
		}

		// Save the transaction
		if _, err := a.Db.NewInsert().Model(data).Exec(ctx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetByID implements entities.TransactionRepository.
func (a *transactionRepository) GetByID(ctx context.Context, id string) (*entities.TransactionModel, error) {
	idAsUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var data entities.TransactionModel
	err = a.Db.NewSelect().Model(&data).Where("id = ?", idAsUuid).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func New(db *bun.DB) entities.TransactionRepository {
	return &transactionRepository{
		Db: db,
	}
}
