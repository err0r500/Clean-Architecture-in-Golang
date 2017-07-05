package iTesting

import (
	"errors"
	"fmt"
)

type User struct {
	id   int
	name string
}

type Order struct {
	id     int
	userID int
}

type orderReader interface {
	GetOrder(id int) (*Order, error)
	GetUser(id int) (*User, error)
}

func CheckOrder(oR orderReader, orderID int) error {
	order, err := oR.GetOrder(orderID)
	if err != nil || order == nil {
		return errors.New("can't get Order")
	}

	user, err := oR.GetUser(order.userID)
	if err != nil {
		return errors.New("can't get User")
	}

	fmt.Printf("Order %d belongs to user %s !", order.id, user.name)
	return nil
}
