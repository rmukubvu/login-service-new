package service

import (
	"golang.org/x/crypto/bcrypt"
	"login-service/model"
	"login-service/repo"
)

func AddUser(user model.User) error {
	hash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return repo.Auth(user)
}

func Login(username, password string) (model.UserLoginResponse, error) {
	user := model.User{}
	errorMessage := "invalid username/password combination"

	user, err := repo.LoginWith(username)
	if err != nil {
		return model.UserLoginResponse{IsError: true, ErrorMessage: err.Error()}, err
	}

	if checkPasswordHash(password, user.Password) {
		errorMessage = ""
	}
	return model.UserLoginResponse{
		Id:           user.Id.Hex(),
		UserName:     user.UserName,
		IsError:      len(errorMessage) > 0,
		ErrorMessage: errorMessage,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TearDown() {
	repo.CloseDb()
}
