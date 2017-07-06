package interfaces

import uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling"

type EvilInterface struct {
	GetOrderOutput GetOrderReturn
	GetUserOutput  GetUserReturn
}

type GetOrderReturn struct {
	Order *uc.Order
	Err   error
}
type GetUserReturn struct {
	User *uc.User
	Err  error
}

func (tI EvilInterface) GetOrder(id int) (*uc.Order, error) {
	return tI.GetOrderOutput.Order, tI.GetOrderOutput.Err
}

func (tI EvilInterface) GetUser(id int) (*uc.User, error) {
	return tI.GetUserOutput.User, tI.GetUserOutput.Err
}
