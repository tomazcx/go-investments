package accountuc

import (
	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
)

type UpdateAccountUC struct {
	repo repository.IAccountRepository
}

func (uc *UpdateAccountUC) Execute(input accountio.UpdateAccountInput) error {
	acc, err := uc.repo.FindById(input.ID)

	if err != nil {
		return err
	}

	acc.Forename = input.Forename
	acc.Surname = input.Surname

	return uc.repo.Update(acc)
}
