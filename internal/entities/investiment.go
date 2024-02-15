package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidInvestmentType = errors.New("Tipo de investimento inválido")
var ErrInvalidInvestmentDate = errors.New("Data de início / final inválida")
var ErrInvalidInvestmentOwner = errors.New("Conta inválida")

const (
	SavingAccount = iota
)

type Investment struct {
	ID        uuid.UUID
	Owner     *Account
	Type      uint8
	StartDate time.Time
	EndDate   time.Time
}

func NewInvestment(owner *Account, investmentType uint8, startDate time.Time, endDate time.Time) (*Investment, error) {
	if investmentType != SavingAccount {
		return nil, ErrInvalidInvestmentType
	}

	if startDate.Before(endDate) || startDate.Equal(endDate) {
		return nil, ErrInvalidInvestmentDate
	}

	if owner == nil {
		return nil, ErrInvalidInvestmentOwner
	}

	return &Investment{
		ID: uuid.New(),
		Type: investmentType,
		Owner: owner,
		StartDate: startDate,
		EndDate: endDate,
	}, nil
}
