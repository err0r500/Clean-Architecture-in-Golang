package iTesting

type niceInterface struct{}

func (tI niceInterface) getOrder(id int) (*order, error) {
	return &order{}, nil
}

func (tI niceInterface) getUser(id int) (*user, error) {
	return &user{}, nil
}
