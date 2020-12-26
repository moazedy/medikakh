package logic

import "errors"

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
