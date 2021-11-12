package _3_Facade

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("el usuario no está logueado")

type ValidatorToken struct {
	token string
}

func NewValidatorToken(t string) ValidatorToken {
	return ValidatorToken{token: t}
}

func (myValidatorToken *ValidatorToken) Validate() error {
	if myValidatorToken.token != "token-valido" {
		return ErrTokenNotValid
	}
	fmt.Println("token válido")
	return nil
}