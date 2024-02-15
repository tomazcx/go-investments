package accountuc

import (
	"time"

	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
)

type DepositUC struct {
	repo repository.IAccountRepository
}

func (uc *DepositUC) Execute(input accountio.DepositInput) (*accountio.DepositOutput, error) {
	acc, err := uc.repo.FindById(input.ID)	
	if err != nil {
		return nil, err
	}

	newBalance := acc.Balance + input.Ammount

	err = uc.repo.UpdateBalance(input.ID, newBalance)
	if err != nil {
		return nil, err
	}

	return &accountio.DepositOutput{
		NewBalance: newBalance,
		Timestamp: time.Now(),
	}, nil
}
