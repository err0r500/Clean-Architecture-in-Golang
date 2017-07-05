package interfaces

import uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling"

type EvilInterface struct {
	GetOrderOutput int
	GetUserOutput  int
}

func (tI EvilInterface) GetOrder(id int) (*uc.Order, error) {
	toReturn := []*uc.Order{
		&uc.Order{},
		nil,
	}
	return toReturn[tI.GetOrderOutput], nil
}

func (tI EvilInterface) GetUser(id int) (*uc.User, error) {
	return &uc.User{}, nil
}
