package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TransactionCreatePayload struct {
	SourceAccountID      int64   `json:"source_account_id" validate:"number,required"`
	DestinationAccountID int64   `json:"destination_account_id" validate:"number,required"`
	Amount               float64 `json:"amount" validate:"number,required"`
}

type TransactionModel struct {
	bun.BaseModel `bun:"table:transactions,alias:t" swaggerignore:"true"`

	ID                   uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	SourceAccountID      int64     `bun:"source_account_id,notnull" json:"source_account_id" validate:"number,required"`
	DestinationAccountID int64     `bun:"destination_account_id,notnull" json:"destination_account_id" validate:"number,required"`
	Amount               float64   `bun:"amount,notnull" json:"amount"`
	CreatedAt            time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"createdAt"`
}

type TransactionResponse struct {
	Status  bool              `json:"status"`
	Data    *TransactionModel `json:"data"`
	Message string            `json:"message,omitempty"`
	Errors  []string          `json:"errors,omitempty"`
}

type TransactionService interface {
	Create(ctx context.Context, input TransactionCreatePayload) (*TransactionModel, error)
	GetByID(ctx context.Context, id string) (*TransactionModel, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, input TransactionCreatePayload) (*TransactionModel, error)
	GetByID(ctx context.Context, id string) (*TransactionModel, error)
}
