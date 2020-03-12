package basket

import (
	"github.com/farwydi/demo-cleanarh/domain"
	"github.com/farwydi/demo-cleanarh/gateway"
	"github.com/farwydi/demo-cleanarh/usecase"
)

func New(basket gateway.BasketGateway) usecase.BasketUseCase {
	return &basketUseCase{
		BasketGateway: basket,
	}
}

type basketUseCase struct {
	BasketGateway gateway.BasketGateway
}

func (uc *basketUseCase) AddItem(user *domain.User, item *domain.Item) error {
	return nil
}

func (uc *basketUseCase) RemoveItem(user *domain.User, item *domain.Item) error {
	uc.BasketGateway.RemoveItemByIdFromUserId(user.ID, item.ID)
	return nil
}

func (uc *basketUseCase) Cleanup(user *domain.User) error {
	return nil
}

func (uc *basketUseCase) Buy(user *domain.User) error {
	return nil
}
