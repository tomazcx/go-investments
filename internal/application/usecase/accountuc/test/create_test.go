package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/go-investments/internal/application/io/accountio"
	"github.com/tomazcx/go-investments/internal/application/protocol/repository/mocks"
	"github.com/tomazcx/go-investments/internal/application/usecase/accountuc"
)

func TestCreateAccountUC_Success(t *testing.T){
	accRepo := &mocks.AccountRepositoryMock{}
	input := accountio.CreateAccountInput{
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
	}
	createdAt := time.Now()
	accRepo.On("EmailExists", input.Email).Return(false, (error)(nil))
	accRepo.On("DocumentExists", input.DocumentID).Return(false, (error)(nil))
	accRepo.On("Create", mock.Anything).Return(createdAt, (error)(nil))

	uc := accountuc.NewCreateAccountUC(accRepo)
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, output.Account.Forename, input.Forename)
	assert.Equal(t, output.Account.Surname, input.Surname)
	assert.Equal(t, output.Account.Type, input.Type)
	assert.Equal(t, output.Account.DocumentID, input.DocumentID)
	assert.Equal(t, output.Account.Email, input.Email)
	assert.NotNil(t, output.Account.Password)
	assert.Equal(t, output.CreatedAt, createdAt)

	accRepo.AssertCalled(t, "EmailExists", input.Email)
	accRepo.AssertCalled(t, "DocumentExists", input.DocumentID)
	accRepo.AssertCalled(t, "Create", mock.Anything)
}

func TestCreateAccountUC_Fail_Duplicated_Email(t *testing.T){
	accRepo := &mocks.AccountRepositoryMock{}
	input := accountio.CreateAccountInput{
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
	}
	accRepo.On("EmailExists", input.Email).Return(true, (error)(nil))

	uc := accountuc.NewCreateAccountUC(accRepo)
	_, err := uc.Execute(input)

	assert.Equal(t, err, accountuc.ErrCreateAccEmailAlreadyRegistered)

	accRepo.AssertCalled(t, "EmailExists", input.Email)
	accRepo.AssertNotCalled(t, "Create")
}

func TestCreateAccountUC_Fail_Duplicated_Document(t *testing.T){
	accRepo := &mocks.AccountRepositoryMock{}
	input := accountio.CreateAccountInput{
		Forename:   "John",
		Surname:    "Doe",
		Email:      "john@gmail.com",
		Type:       0,
		Password:   "12345",
		DocumentID: "18381387335",
	}
	accRepo.On("EmailExists", input.Email).Return(false, (error)(nil))
	accRepo.On("DocumentExists", input.DocumentID).Return(true, (error)(nil))

	uc := accountuc.NewCreateAccountUC(accRepo)
	_, err := uc.Execute(input)

	assert.Equal(t, err, accountuc.ErrCreateAccDocumentAlreadyRegistered)

	accRepo.AssertCalled(t, "EmailExists", input.Email)
	accRepo.AssertCalled(t, "DocumentExists", input.DocumentID)
	accRepo.AssertNotCalled(t, "Create")
}
