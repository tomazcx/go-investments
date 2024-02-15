package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidInvestmentType = errors.New("Tipo de investimento inválido")
var ErrInvalidInvestmentDate = errors.New("Data de início / final inválida")
var ErrInvalidInvestmentOwner = errors.New("Conta inválida")
var ErrInvalidInvestmentAmount = errors.New("Valor inválido. Deve ser no mínimo R$100,00")

const (
	SavingAccount = iota
)

type Investment struct {
	ID            uuid.UUID
	Owner         *Account
	Type          uint8
	StartDate     time.Time
	EndDate       time.Time
	InitialAmount float64
}

func NewInvestment(owner *Account, investmentType uint8, startDate time.Time, endDate time.Time, initialAmount float64) (*Investment, error) {
	if investmentType != SavingAccount {
		return nil, ErrInvalidInvestmentType
	}

	if endDate.Before(startDate) || startDate.Equal(endDate) {
		return nil, ErrInvalidInvestmentDate
	}

	if owner == nil {
		return nil, ErrInvalidInvestmentOwner
	}

	if initialAmount < 100.00 {
		return nil, ErrInvalidInvestmentAmount
	}

	return &Investment{
		ID:        uuid.New(),
		Type:      investmentType,
		Owner:     owner,
		StartDate: startDate,
		EndDate:   endDate,
		InitialAmount: initialAmount,
	}, nil
}
