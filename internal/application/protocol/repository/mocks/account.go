package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/tomazcx/go-investments/internal/entities"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Exists(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}


func (m *AccountRepositoryMock) FindById(id string) (*entities.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Account), args.Error(1)
}


func (m *AccountRepositoryMock) DocumentExists(document string) (bool, error) {
	args := m.Called(document)
	return args.Get(0).(bool), args.Error(1)
}

func (m *AccountRepositoryMock) EmailExists(email string) (bool, error) {
	args := m.Called(email)
	return args.Get(0).(bool), args.Error(1)
}

func (m *AccountRepositoryMock) FindByEmail(email string) (*entities.Account, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.Account), args.Error(1)
}

func (m *AccountRepositoryMock) Create(account *entities.Account) (time.Time, error) {
	args := m.Called(account)
	return args.Get(0).(time.Time), args.Error(1)
}

func (m *AccountRepositoryMock) Update(account *entities.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountRepositoryMock) UpdateBalance(id string, newBalance float32) error {
	args := m.Called(id, newBalance)
	return args.Error(0)
}

func (m *AccountRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
