// Code generated by ifacemaker

package gateway

import (
	"github.com/farwydi/demo-cleanarh/domain"
)

// BasketGateway ...
type BasketGateway interface {
	RemoveItemByIdFromUserId(userID domain.ID, itemID domain.ID)
}