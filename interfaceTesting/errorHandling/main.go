package iTesting

import "fmt"

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
	order, _ := oR.GetOrder(orderID)

	user, _ := oR.GetUser(order.UserID)

	fmt.Printf("Order %d belongs to user %s !\n", order.ID, user.Name)
	return nil
}
