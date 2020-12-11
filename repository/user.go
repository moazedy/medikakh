package repository

import (
	"medikakh/domain/models"
	"medikakh/repository/queries"

	"github.com/couchbase/gocb/v2"
)

type UserRepo interface {
	Register(user models.User) error
	ReadUserById(userId string) (*models.User, error)
	ReadUserByUsername(username string) (*models.User, error)
	UpdateUser(newUser models.User) error
	DeleteUser(Id string) error
	GetUserPassword(Id string) (*string, error)
	GetUserRole(Id string) (*string, error)
	GetUserIdByUsername(username string) (*string, error)
	IsUsernameExists(username string) (*bool, error)
}

type user struct {
	session *gocb.Cluster
}

func NewUserRpo(session *gocb.Cluster) UserRepo {
	u := new(user)
	u.session = session
	return u
}

func (u *user) Register(user models.User) error {
	_, err := u.session.Query(
		queries.RegisterUserQuery,
		&gocb.QueryOptions{
			NamedParameters: map[string]interface{}{
				"id":         user.Id,
				"username":   user.Username,
				"password":   user.Password,
				"role":       user.Role,
				"created_at": user.CreatedAt,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) ReadUserById(userId string) (*models.User, error) {
	res, err := u.session.Query(
		queries.ReadUserByIdQuery,
		&gocb.QueryOptions{
			PositionalParameters: []interface{}{userId},
		},
	)
	if err != nil {
		return errors.New("error on getting data from db")
	}

	var newUser models.User
	for res.Next() {
		err = res.Row(&newUser)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("user does not exist !")
			}
			return nil, err
		}
	}
	return newUser, nil
}

func (u *user) ReadUserByUsername(username string) (*models.User, error) {
	res, err := u.session.Query(
		queries.ReadUserByUsername,
		&gocb.QueryOptions{PositionalParameters: []interface{}{username}},
	)
	if err != nil {
		return nil, errors.New("error on reading data form db")
	}

	var newUser models.User
	for res.Next() {
		err = res.Row(&newUser)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("user does not exist !")
			}
			return nil, err
		}
	}
	return newUser, nil

}

func (u *user) DeleteUser(Id string) error {
	_, err := u.session.Query(
		queries.DeleteUserQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{Id}},
	)
	if err != nil {
		return errors.New("error in interactin with database")
	}

	return nil
}

func (u *user) GetUserPassword(Id string) (*string, error) {
	res, err := u.session.Query(
		queries.GetUserPasswordQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{Id}},
	)
	if err != nil {
		return nil, errors.New("error on reading data from db")
	}

	var pass string
	for res.Next() {
		err = res.Row(&pass)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("user does not exist")
			}
			return nil, err
		}
	}

	return &pass, nil
}

func (u *user) GetUserRole(Id string) (*string, error) {
	res, err := u.session.Query(
		queries.GetUserRole,
		&gocb.QueryOptions{PositionalParameters: []interface{}{Id}},
	)
	if err != nil {
		return nil, errors.New("error on reading data from db")
	}

	var role string
	for res.Next() {
		err = res.Row(&role)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("user does not exist")
			}
			return nil, err
		}
	}

	return &role, nil
}

func (u *user) GetUserIdByUsername(username string) (*string, error) {
	res, err := u.session.Query(
		queries.GetUserIdByUsernameQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{username}},
	)
	if err != nil {
		return nil, errors.New("error on reading data from db")
	}

	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("user does not exist")
			}
			return nil, err
		}
	}

	return &id, nil

}

func (u *user) UpdateUser(newUser models.User) error {
	_, err := u.session.Query(
		queries.UpdateUserQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{
			newUser.Username,
			newUser.Password,
			newUser.Role,
			newUser.CreatedAt,
		}},
	)
	if err != nil {
		return errors.New("error on working wiht db")
	}

	return nil
}

func (u *user) IsUsernameExists(username string) (*bool, error) {
	res, err := u.session.Query(
		queries.IsUsernameExistsQuery,
		&gocb.QueryOptions{PositinalParameters: []interface{}{username}},
	)
	if err != nil {
		return nil, errors.New("error on serching for specific username")
	}

	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("user does not exist")
			}
			return nil, err
		}
	}

	if id != "" || id != nil {
		return true, nil
	}

	return false, nil
}
