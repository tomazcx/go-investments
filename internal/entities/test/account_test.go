package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomazcx/go-investments/internal/entities"
)

func TestCreateNewAccount_Success(t *testing.T){
	forename := "John"
	surname := "Doe"
	email := "john@gmail.com"
	accType := 0
	password := "12345"
	documentId := "18381387335"

	acc, err := entities.NewAccount(forename, surname, email, uint8(accType), documentId, password)
	assert.Nil(t, err)
	assert.Equal(t, acc.Forename, forename)
	assert.Equal(t, acc.Surname, surname)
	assert.Equal(t, acc.Email, email)
	assert.Equal(t, acc.Type, uint8(accType))
	assert.True(t, acc.IsPasswordValid(password))
	assert.NotNil(t, acc.ID)
	assert.Equal(t, acc.GetName(), forename + " " + surname)
}

func TestCreateNewAccount_Invalid_Type(t *testing.T){
	forename := "John"
	surname := "Doe"
	email := "john@gmail.com"
	accType := 2
	password := "12345"
	documentId := "18381387335"

	_, err := entities.NewAccount(forename, surname, email, uint8(accType), password, documentId)
	assert.Equal(t, err, entities.ErrInvalidAccType)

	accType = 1
	_, err = entities.NewAccount(forename, surname, email, uint8(accType), password, documentId)
	assert.Equal(t, err, entities.ErrInvalidAccDocument)

	accType = 0
	documentId = "18381387335324"
	_, err = entities.NewAccount(forename, surname, email, uint8(accType), password, documentId)
	assert.Equal(t, err, entities.ErrInvalidAccDocument)
}
