package store

import (
	"matester/pkg/db"
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

type AuthValidator interface {
	IsAuthorised(login string, authRow *db.AuthRow) bool
}

type AuthValidatorImpl struct {
}

func NewAuthValidator() AuthValidator {
	return &AuthValidatorImpl {}
}

func (v *AuthValidatorImpl) IsAuthorised(login string, authRow *db.AuthRow) bool {
	got := GetMD5Hash(login + authRow.Salt)
	fmt.Printf("Check %s ? %s", got, authRow.Pass)

	return string(got) == authRow.Pass
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
 }