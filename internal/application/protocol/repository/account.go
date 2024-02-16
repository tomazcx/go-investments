package repository

import (
	"time"

	"github.com/tomazcx/go-investments/internal/entities"
)

type IAccountRepository interface{
	Exists(string) (bool, error)
	FindById(string) (*entities.Account, error)
	EmailExists(string) (bool, error)
	DocumentExists(string) (bool, error)
	FindByEmail(string) (*entities.Account, error)
	Create(*entities.Account) (time.Time, error)
	Update(*entities.Account) error
	UpdateBalance(string, float32) error
	Delete(string) error
}


