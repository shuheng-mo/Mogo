package service

import (
	"errors"
	"github.com/shuheng-mo/Mogo/user/domain/model"
	"github.com/shuheng-mo/Mogo/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOK bool, err error)
}

// NewUserDataService New instance
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserReposiory: userRepository}
}

type UserDataService struct {
	UserReposiory repository.IUserRepository
}

// GeneratePassword generated encrypted user password
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword validate hashed user password
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("Compare hashed password error")
	}
	return true, nil
}

// AddUser add new user
func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserReposiory.CreateUser(user)
}

// DeleteUser delete a user
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserReposiory.DeleteUserByID(userID)
}

// UpdateUser update user
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	// check whether the password has been updated by the user
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	// write to logger
	return u.UserReposiory.UpdateUser(user)
}

// FindUserByName find the user by name
func (u *UserDataService) FindUserByName(userName string) (user *model.User, err error) {
	return u.UserReposiory.FindUserByName(userName)
}

// CheckPwd check the hashed user password
func (u *UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserReposiory.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
