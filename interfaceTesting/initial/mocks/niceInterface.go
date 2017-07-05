package iTesting

import uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/initial"

type NiceInterface struct{}

func (tI NiceInterface) GetOrder(id int) (*uc.Order, error) {
	return &uc.Order{}, nil
}

func (tI NiceInterface) GetUser(id int) (*uc.User, error) {
	return &uc.User{}, nil
}
