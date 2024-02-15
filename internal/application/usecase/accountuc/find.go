package accountuc

import (
	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
)

type FindAccountUC struct {
	repo repository.IAccountRepository
}

func (uc *FindAccountUC) Execute(id string) (*accountio.FindAccountOutput, error) {
	acc, err := uc.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return &accountio.FindAccountOutput{
		Account: acc,
	}, nil
}
