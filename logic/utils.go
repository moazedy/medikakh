package logic

import (
	"errors"
	"time"
)

func checkForRoleStatmentCorrectness(role string) bool {
	switch role {
	case "bronze":
		return true
	case "silver":
		return true
	case "gold":
		return true
	default:
		return false
	}
}

func checkUsernameValueValidation(username string) error {
	if username == "" {
		return errors.New("username can not be empty")
	}

	if len(username) > 30 {
		return errors.New("too long username, username should be less than 30 characters")
	}

	if len(username) < 2 {
		return errors.New("too short username")
	}

	return nil
}

func checkPasswordValueValidation(pass string) error {
	if pass == "" {
		return errors.New("password can't be empty")
	}

	if len(pass) > 30 {
		return errors.New("too long password, password should be less than 30 characters")
	}

	if len(pass) < 4 {
		return errors.New("too short password, it should be at least 4 characters")
	}

	return nil
}

// GetRolePeriod returns number of every role's day of trail
func GetRolePeriod(role string) (int, error) {
	switch role {
	case "bronze":
		return 30, nil
	case "silver":
		return 90, nil
	case "gold":
		return 180, nil
	default:
		return 0, errors.New("role value is incorrect")
	}
}

func CheckingTimeExpiration(role string, createdAt time.Time) (*bool, error) {
	days, err := GetRolePeriod(role)
	if err != nil {
		return nil, err
	}

	var exp bool
	hours := time.Duration(days * 24)
	expTime := createdAt.Add(time.Hour * hours)
	if expTime.After(time.Now()) {
		return &exp, nil
	}

	exp = true
	return &exp, nil

}
