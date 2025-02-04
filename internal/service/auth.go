package service

type UserRepository interface {
}

type Auth struct {
	userRepo UserRepository
}

func NewAuth(userRepo UserRepository) *Auth {
	return &Auth{
		userRepo: userRepo,
	}
}

func (a *Auth) Login() {

}

func (a *Auth) Register() {

}
