package repository

import (
	"context"

	"github.com/ariefsn/intrans/entities"
	"github.com/uptrace/bun"
)

type accountRepository struct {
	Db *bun.DB
}

// Create implements entities.AccountRepository.
func (a *accountRepository) Create(ctx context.Context, input entities.AccountCreatePayload) (*entities.AccountModel, error) {
	data := &entities.AccountModel{
		ID:      input.ID,
		Balance: input.InitialBalance,
	}

	_, err := a.Db.NewInsert().Model(data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetByID implements entities.AccountRepository.
func (a *accountRepository) GetByID(ctx context.Context, id int64) (*entities.AccountModel, error) {
	var data entities.AccountModel
	err := a.Db.NewSelect().Model(&data).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func New(db *bun.DB) entities.AccountRepository {
	return &accountRepository{
		Db: db,
	}
}
