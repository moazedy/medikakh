package logic

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medikakh/domain/models"
	"medikakh/repository"
	"medikakh/service/payment"
	"time"
)

type UserLogic interface {
	Register(username, password, role string) error
	ReadUser(username string) (*models.User, error)
}

type user struct {
	repo repository.UserRepo
}

func NewUserLogic(repo repository.UserRepo) UserLogic {
	u := new(user)
	u.repo = repo
	return u
}

func (u *user) Register(username, password, role string) error {

	roleCorrectness := checkForRoleStatmentCorrectness(role)
	if !roleCorrectness {
		return errors.New("role statment is invalid")
	}

	err := checkUsernameValueValidation(username)
	if err != nil {
		return err
	}

	err = checkPasswordValueValidation(password)
	if err != nil {
		return err
	}

	paymentResult := payment.RedirectToPay(role)
	if !paymentResult {
		return errors.New("error in payment")
	}

	userExistance, err := u.repo.IsUsernameExists(username)
	if *userExistance {
		return errors.New("username alredy exists")
	}

	var newUser models.User
	newUser.Id = uuid.New()
	newUser.Username = username
	newUser.Role = role
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPass)
	newUser.CreatedAt = time.Now()

	err = u.repo.Register(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *user) ReadUser(username string) (*models.User, error) {
	err := checkUsernameValueValidation(username)
	if err != nil {
		return nil, err
	}

	userExistance, err := u.repo.IsUsernameExists(username)
	if err != nil {
		return nil, err
	}
	if !*userExistance {
		return nil, errors.New("username does not exist !")
	}

	user, err := u.repo.ReadUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
