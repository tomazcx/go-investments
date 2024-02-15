package accountuc

import (
	"errors"

	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
	"github.com/tomazcx/go-investments/internal/entities"
)

var ErrCreatAccEmailAlreadyRegistered = errors.New("Email already registered")

type CreateAccountUC struct {
	repo repository.IAccountRepository
}

func (uc *CreateAccountUC) Execute(input accountio.CreateAccountInput) (*accountio.CreateAccountOutput, error) {
	emailExists, err := uc.repo.EmailExists(input.Email)	

	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, ErrCreatAccEmailAlreadyRegistered
	}

	acc, err := entities.NewAccount(
		input.Forename,
		input.Surname,
		input.Email,
		input.Type,
		input.DocumentId,
		input.Password,
		input.InitialBalance,
	)

	if err != nil {
		return nil, err
	}

	createdAt, err := uc.repo.Create(acc)

	return &accountio.CreateAccountOutput{
		Account: acc,
		CreatedAt: createdAt,
	}, nil
}
