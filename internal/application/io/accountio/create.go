package accountio

import (
	"time"

	"github.com/tomazcx/go-investments/internal/entities"
)

type CreateAccountInput struct {
	Forename       string
	Surname        string
	Email          string
	InitialBalance float32
	Password       string
	Type           uint8
	DocumentID     string
}

type CreateAccountOutput struct {
	Account *entities.Account
	CreatedAt time.Time
}
