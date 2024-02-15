package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/go-investments/internal/entities"
)

func TestCreateInvestment_Success(t *testing.T) {
	owner := &entities.Account{
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
	}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)
	initialAmount := 100.50
	investmentType := 0

	investment, err := entities.NewInvestment(owner, uint8(investmentType), startDate, endDate, initialAmount)
	assert.Nil(t, err)
	assert.Equal(t, investment.Owner, owner)
	assert.Equal(t, investment.Type, uint8(investmentType))
	assert.Equal(t, investment.StartDate, startDate)
	assert.Equal(t, investment.EndDate, endDate)
	assert.Equal(t, investment.InitialAmount, initialAmount)
}

func TestCreateInvestment_Fail(t *testing.T) {
	owner := &entities.Account{
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
	}
	startDate := time.Now()
	endDate := startDate.AddDate(0, 2, 0)
	initialAmount := 100.50
	investmentType := 0

	_, err := entities.NewInvestment(nil, uint8(investmentType), startDate, endDate, initialAmount)
	assert.Equal(t, err, entities.ErrInvalidInvestmentOwner)

	endDate = endDate.AddDate(0, -5, 0)
	_, err = entities.NewInvestment(owner, uint8(investmentType), startDate, endDate, initialAmount)
	assert.Equal(t, err, entities.ErrInvalidInvestmentDate)

	endDate = endDate.AddDate(0, 10, 0)
	_, err = entities.NewInvestment(owner, 1, startDate, endDate, initialAmount)
	assert.Equal(t, err, entities.ErrInvalidInvestmentType)

	initialAmount = 99.99
	_, err = entities.NewInvestment(owner, uint8(investmentType), startDate, endDate, initialAmount)
	assert.Equal(t, err, entities.ErrInvalidInvestmentAmount)
}
