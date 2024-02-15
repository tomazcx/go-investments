package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/go-investments/internal/application/usecase/accountuc"
	"github.com/tomazcx/go-investments/internal/entities"
)

func TestWithdraw_Success(t *testing.T){
	accRepo := &mocks.AccountRepositoryMock{}
	id := uuid.New()
	input := accountio.WithdrawInput{
		ID: id.String() ,
		Ammount: 29.90,
	}
	acc := &entities.Account{
		ID: id,
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
		Balance: 40.90,
	}
	newBalance := acc.Balance - input.Ammount
	accRepo.On("FindById", input.ID).Return(acc, (error)(nil))
	accRepo.On("UpdateBalance", input.ID,newBalance).Return((error)(nil))

	uc := accountuc.NewWithdrawUC(accRepo)
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, output.NewBalance, newBalance)

	accRepo.AssertCalled(t,"FindById", input.ID)
	accRepo.AssertCalled(t,"UpdateBalance", input.ID, newBalance)
}

func TestWithdraw_Fail(t *testing.T){
	accRepo := &mocks.AccountRepositoryMock{}
	input := accountio.WithdrawInput{
		ID: uuid.New().String(),
		Ammount: 29.90,
	}
	acc := &entities.Account{
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
		Balance: 19.90,
	}
	accRepo.On("FindById", input.ID).Return(acc, (error)(nil))

	uc := accountuc.NewWithdrawUC(accRepo)
	_, err := uc.Execute(input)

	assert.Equal(t, err, accountuc.ErrWithdrawNotEnoughtBalance)

	accRepo.AssertCalled(t,"FindById", input.ID)
	accRepo.AssertNotCalled(t,"UpdateBalance")
}
