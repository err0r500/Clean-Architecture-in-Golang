package iTesting

import (
	"errors"
	"fmt"
)

//User the guy who placed an order
type user struct {
	id   int
	name string
}

//Order content on an Order
type order struct {
	id     int
	userID int
}

type orderReader interface {
	getOrder(id int) (*order, error)
	getUser(id int) (*user, error)
}

func checkOrder(oR orderReader, orderID int) error {
	order, err := oR.getOrder(orderID)
	if err != nil {
		return errors.New("can't get Order")
	}

	user, err := oR.getUser(order.userID)
	if err != nil {
		return errors.New("can't get User")
	}

	fmt.Printf("Order %d belongs to user %s !", order.id, user.name)
	return nil
}
