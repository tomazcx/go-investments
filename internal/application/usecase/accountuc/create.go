package accountuc

import (
	"errors"

	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository"
	"github.com/tomazcx/go-investments/internal/entities"
)

var ErrCreateAccEmailAlreadyRegistered = errors.New("Email já registrado")
var ErrCreateAccDocumentAlreadyRegistered = errors.New("Email já registrado")

type CreateAccountUC struct {
	repo repository.IAccountRepository
}

func (uc *CreateAccountUC) Execute(input accountio.CreateAccountInput) (*accountio.CreateAccountOutput, error) {
	emailExists, err := uc.repo.EmailExists(input.Email)

	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, ErrCreateAccEmailAlreadyRegistered
	}

	documentExists, err := uc.repo.DocumentExists(input.DocumentID)

	if err != nil {
		return nil, err
	}

	if documentExists {
		return nil, ErrCreateAccDocumentAlreadyRegistered 
	}

	acc, err := entities.NewAccount(
		input.Forename,
		input.Surname,
		input.Email,
		input.Type,
		input.DocumentID,
		input.Password,
		input.InitialBalance,
	)

	if err != nil {
		return nil, err
	}

	createdAt, err := uc.repo.Create(acc)

	return &accountio.CreateAccountOutput{
		Account:   acc,
		CreatedAt: createdAt,
	}, nil
}

func NewCreateAccountUC(repo repository.IAccountRepository) *CreateAccountUC {
	return &CreateAccountUC{
		repo: repo,
	}
}
