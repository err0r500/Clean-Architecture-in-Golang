package iTesting

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Order struct {
	ID     int
	UserID int
}

type orderReader interface {
	GetOrder(id int) (*Order, error)
	GetUser(id int) (*User, error)
}

func CheckOrder(oR orderReader, orderID int) error {
	order, err := oR.GetOrder(orderID)
	if err != nil || order == nil {
		return errors.New("unable to retreive the order")
	}
	if order.ID != orderID || order.UserID == 0 {
		return errors.New("the order returned is wrong")
	}

	user, err := oR.GetUser(order.UserID)
	if err != nil || user == nil {
		return errors.New("unable to retreive the user")
	}
	if order.UserID != user.ID {
		return errors.New("the user returned is wrong")
	}
	fmt.Printf("Order %d belongs to user %s !\n", order.ID, user.Name)
	return nil
}
