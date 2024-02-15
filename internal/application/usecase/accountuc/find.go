package accountuc

import (
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
	"github.com/tomazcx/go-investments/internal/entities"
)

type FindAccountUC struct {
	repo repository.IAccountRepository
}

func (uc *FindAccountUC) Execute(id string) (*entities.Account, error) {
	return uc.repo.FindById(id)
}

func NewFindAccountUC(repo repository.IAccountRepository) *FindAccountUC {
	return &FindAccountUC{
		repo: repo,
	}
}
