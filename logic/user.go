package logic

import (
	"errors"
	"fmt"
	"medikakh/application/utils"
	"medikakh/domain/constants"
	"medikakh/domain/models"
	"medikakh/repository"
	"medikakh/service/authorization"
	"medikakh/service/payment"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserLogic interface {
	Register(username, password, role, email string) error
	ReadUser(userRole, userId, username string) (*models.User, error)
	RevivalAcount(c *gin.Context, username, role, email string) error
	IsUserExists(userRole, username string) error
	GetUserRole(userRole, username string) (*string, error)
	GetUserId(userRole, username string) (*string, error)
	GetPassword(userRole, RequesterUserId, id string) (*string, error)
	UpdateUser(userRole, userId string, us models.UserUpdate) error
}

type user struct {
	repo repository.UserRepo
}

func NewUserLogic(repo repository.UserRepo) UserLogic {
	u := new(user)
	u.repo = repo
	return u
}

func (u *user) Register(username, password, role, email string) error {

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

func (u *user) ReadUser(userRole, userId, username string) (*models.User, error) {
	err := checkUsernameValueValidation(username)
	if err != nil {
		return nil, err
	}
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
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

	if userRole != constants.AdminRoleObject || userRole != constants.SystemRoleObject {
		if userId != user.Id.String() {
			return nil, errors.New("no permissions on reading this user")
		}
	}

	return user, nil
}

// RevivalAcount is been used when some acount is being expired and user wants to revivals it
func (u *user) RevivalAcount(c *gin.Context, username, role, email string) error {
	err := checkUsernameValueValidation(username)
	if err != nil {
		return err
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(role, constants.UserObject, constants.UpdateAction)
	if !ok {
		return errors.New("premission denied")
	}

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

	Ok, err := payment.RedirectToPay(c, role, email)
	if !Ok {
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

func (u *user) IsUserExists(userRole, username string) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.ReadAction)
	if !ok {
		return errors.New("premission denied")
	}

	userExistance, err := u.repo.IsUsernameExists(username)
	if err != nil {
		return err
	}
	if !userExistance {
		return errors.New("user does not exist")
	}

	return nil

}

func (u *user) GetUserRole(userRole, username string) (*string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	err := utils.CheckUsernameValueValidation(username)
	if err != nil {
		return nil, err
	}

	userId, err := u.repo.GetUserIdByUsername(username)
	if err != nil {
		return nil, err
	}

	Role, err := u.repo.GetUserRole(*userId)
	if err != nil {
		return nil, err
	}

	return Role, nil

}

func (u *user) GetUserId(userRole, username string) (*string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	name, err := u.repo.GetUserIdByUsername(username)
	if err != nil {
		return nil, err
	}

	return name, nil
}

func (u *user) GetPassword(userRole, requesterUserId, id string) (*string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	if userRole != constants.AdminRoleObject || userRole != constants.SystemRoleObject {
		if requesterUserId != id {
			return nil, errors.New("permission denied")
		}
	}

	pass, err := u.repo.GetUserPassword(id)
	if err != nil {
		return nil, err
	}

	return pass, nil
}

func (u *user) UpdateUser(userRole, userId string, us models.UserUpdate) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.UpdateAction)
	if !ok {
		return errors.New("premission denied")
	}

	if userRole != "admin" || userId != us.Id.String() {
		return errors.New("access denied")
	}

	oldUser, err := u.repo.ReadUserById(userId)
	if err != nil {
		return err
	}
	var newUser models.User
	newUser.Id = us.Id
	newUser.CreatedAt = oldUser.CreatedAt

	if us.Username != nil {
		existance, err := u.repo.IsUsernameExists(*us.Username)
		if err != nil {
			return err
		}
		if existance {
			return errors.New("username alredy exists")
		}
		newUser.Username = *us.Username
	} else {
		newUser.Username = oldUser.Username
	}

	if us.Password != nil {
		err := utils.CheckPasswordValueValidation(*us.Password)
		if err != nil {
			return err
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(*us.Password), 10)
		if err != nil {
			return err
		}
		newUser.Password = string(hashedPass)
	} else {
		newUser.Password = oldUser.Password
	}

	if us.Email != nil {
		newUser.Email = *us.Email // TODO: email validation
	} else {
		newUser.Email = oldUser.Email
	}

	if us.Role != nil {
		ok := authorization.IsPermissioned(userRole, constants.UserObject, constants.UpdateRoleAction)
		if !ok {
			return errors.New("premission denied")
		}

		newUser.Role = *us.Role
	} else {
		newUser.Role = oldUser.Role
	}

	err = u.repo.UpdateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
