package accountuc

import (
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
)

type DeleteAccountUC struct {
	repo repository.IAccountRepository
}

func (uc *DeleteAccountUC) Execute(id string) error {		
	return uc.repo.Delete(id)
}

func NewDeleteAccountUC(repo repository.IAccountRepository) *DeleteAccountUC{
	return &DeleteAccountUC{
		repo: repo,
	}
}
