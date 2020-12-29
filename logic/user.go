package logic

import (
	"errors"
	"fmt"
	"medikakh/domain/models"
	"medikakh/repository"
	"medikakh/service/payment"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserLogic interface {
	Register(c *gin.Context, username, password, role, email string) error
	ReadUser(username string) (*models.User, error)
	RevivalAcount(c *gin.Context, username, role, email string) error
}

type user struct {
	repo repository.UserRepo
}

func NewUserLogic(repo repository.UserRepo) UserLogic {
	u := new(user)
	u.repo = repo
	return u
}

func (u *user) Register(c *gin.Context, username, password, role, email string) error {

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
	// todo checking email value validation
	paymentResult, err := payment.RedirectToPay(c, role, email)
	if !paymentResult || err != nil {
		return err
	}

	userExistance, err := u.repo.IsUsernameExists(username)
	if err != nil {
		return err
	}
	if userExistance {
		return errors.New("user alredy exists")
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
	if !userExistance {
		return nil, errors.New("user does not exist")
	}

	user, err := u.repo.ReadUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) RevivalAcount(c *gin.Context, username, role, email string) error {
	userExistance, err := u.repo.IsUsernameExists(username)
	if err != nil {
		return err
	}
	if !userExistance {
		return errors.New("user does not exist")
	}

	oldUser, err := u.repo.ReadUserByUsername(username)
	if err != nil {
		return err
	}

	timeExpieration, err := CheckingTimeExpiration(oldUser.Role, oldUser.CreatedAt)
	if err != nil {
		return nil
	}
	if !*timeExpieration {
		return errors.New("the acount has not been expierd yet")
	}

	roleCorrectness := checkForRoleStatmentCorrectness(role)
	if !roleCorrectness {
		return errors.New("role statment is invalid")
	}

	ok, err := payment.RedirectToPay(c, role, email)
	if !ok {
		fmt.Println(err)
		return errors.New("error in payment")
	}

	oldUser.Role = role
	oldUser.CreatedAt = time.Now()
	err = u.repo.UpdateUser(*oldUser)
	if err != nil {
		return errors.New("error on updating user")
	}

	return nil
}
