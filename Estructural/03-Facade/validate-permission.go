package _3_Facade

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("el usuario no está autorizado")

type ValidatorPermission struct {
	userID string
}

func NewValidatorPermission(ID string) ValidatorPermission {
	return ValidatorPermission{userID: ID}
}

func (myValidatorPermission *ValidatorPermission) Validate() error {
	if myValidatorPermission.userID != "user-blog" {
		return ErrPermissionNotValid
	}
	fmt.Println("usuario autorizado")
	return nil
}