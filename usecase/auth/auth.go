package jwt

import (
	"errors"
	"github.com/farwydi/demo-cleanarh/domain"
)

type authUseCase struct {
}

func (gw *authUseCase) Login(fingerprint, login, password string) (*domain.User, error) {
	if len(fingerprint) != 20 {
		return nil, errors.New("fingerprint must be 20 size")
	}

	pwdHash, err := domain.PasswordTakeHash(password)
	if err != nil {
		return nil, err
	}

	panic("implement me")
}

func (gw *authUseCase) Logout(*domain.User) error {
	panic("implement me")
}

func (gw *authUseCase) IsLogin(user *domain.User) (bool, error) {
	panic("implement me")
}
