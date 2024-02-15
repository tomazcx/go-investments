package accountuc

import (
	"errors"
	"time"

	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
)

var ErrWithdrawNotEnoughtBalance = errors.New("Saldo insuficiente para saque")

type WithdrawUC struct {
	repo repository.IAccountRepository
}

func (uc *WithdrawUC) Execute(input accountio.WithdrawInput) (*accountio.WithdrawOutput, error) {
	acc, err := uc.repo.FindById(input.ID)	
	if err != nil {
		return nil, err
	}

	newBalance := acc.Balance - input.Ammount
	if newBalance < 0{
		return nil, ErrWithdrawNotEnoughtBalance	
	}

	err = uc.repo.UpdateBalance(input.ID, newBalance)
	if err != nil {
		return nil, err
	}

	return &accountio.WithdrawOutput{
		NewBalance: newBalance,
		Timestamp: time.Now(),
	}, nil
}

func NewWithdrawUC(repo repository.IAccountRepository) *WithdrawUC {
	return &WithdrawUC{
		repo: repo,
	}
}
