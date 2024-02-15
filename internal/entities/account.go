package entities

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidAccType = errors.New("Tipo inválido")

const (
	PhysicalPerson = iota
	JuridicPerson
)

type Account struct {
	ID       uuid.UUID
	Forename string
	Surname  string
	Email    string
	Type     uint8
	Password string
	DocumentID string //CPF or CNPJ
}

func NewAccount(forename string, surname string, email string, accType uint8, documentId string, password string) (*Account, error) {
	if accType < PhysicalPerson || accType > JuridicPerson {
		return nil, ErrInvalidAccType
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID: uuid.New(),
		Forename: forename,
		Surname: surname,
		Email: email,
		Type: accType,
		DocumentID: documentId,
		Password: string(encryptedPassword),
	}, nil
}

func (a *Account) GetName() string {
	return a.Forename + " " + a.Surname
}

func (a *Account) GetAccountType() (string, error) {
	switch(a.Type){
	case PhysicalPerson:
		return "Pessoa física", nil
	case JuridicPerson:
		return "Pessoa jurídica", nil
	default:
		return "", ErrInvalidAccType  
	}
}
