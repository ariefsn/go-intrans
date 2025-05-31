package entities

import (
	"context"

	"github.com/uptrace/bun"
)

type AccountCreatePayload struct {
	ID             int64   `json:"id" validate:"number,required"`
	InitialBalance float64 `json:"initial_balance" validate:"number,required"`
}

type AccountModel struct {
	bun.BaseModel `bun:"table:accounts,alias:u" swaggerignore:"true"`

	ID      int64   `bun:"id,pk" json:"id"`
	Balance float64 `bun:"balance,notnull" json:"balance"`
}

type AccountResponse struct {
	Status  bool          `json:"status"`
	Data    *AccountModel `json:"data"`
	Message string        `json:"message,omitempty"`
	Errors  []string      `json:"errors,omitempty"`
}

type AccountService interface {
	Create(ctx context.Context, input AccountCreatePayload) (*AccountModel, error)
	GetByID(ctx context.Context, id int64) (*AccountModel, error)
}

type AccountRepository interface {
	Create(ctx context.Context, input AccountCreatePayload) (*AccountModel, error)
	GetByID(ctx context.Context, id int64) (*AccountModel, error)
}
