package service

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"forum/models"
	"forum/repository"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func checkEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func checkString(s string) error {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return errors.New("Empty string")
	}
	return nil
}

func checkCreds(userName, email, password string) error {
	if err := checkEmail(email); err != nil {
		return err
	}
	if err := checkEmail(password); err != nil {
		return err
	}
	if err := checkString(userName); err != nil {
		return err
	}
	return nil
}

func Registration(repos *repository.Repository, userName, email, password string) (models.User, int, error) {
	passHash, err := GeneratePassHash(password)

	newUser := models.User{
		UserName: userName,
		Email:    email,
		PassHash: passHash,
	}
	if err := checkCreds(userName, email, password); err != nil {
		return newUser, http.StatusBadRequest, err
	}

	if err != nil {
		fmt.Println(err.Error())
		return newUser, http.StatusInternalServerError, errors.New("Password Hash Error")
	}

	id, err := repos.Authorization.CreateUser(newUser)
	if err != nil {
		fmt.Println(err.Error())
		return newUser, http.StatusInternalServerError, errors.New("Unable to save to database")

	}
	fmt.Print(id)
	return newUser, http.StatusCreated, nil
}

func GeneratePassHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println(string(hash))
	return string(hash), nil
}

func CompareHashAndPass(passHassFromDb string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passHassFromDb), []byte(password)); err != nil {
		fmt.Println("%w", err)
		return false
	}
	return true
}

func NewToken() (string, error) {
	token, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}
