package iTesting

type evilInterface struct {
	getOrderOutput int
	getUserOutput  int
}

func (tI evilInterface) getOrder(id int) (*order, error) {
	return &order{}, nil
}

func (tI evilInterface) getUser(id int) (*user, error) {
	return &user{}, nil
}
