package logic

type UserLogic interface {
	Register(username, password, role string) error
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

}
