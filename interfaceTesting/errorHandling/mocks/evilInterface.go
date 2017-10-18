package interfaces

import uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling"

type EvilInterface struct {
	GetOrderOutput GetOrderReturn
	GetUserOutput  GetUserReturn
}

type GetOrderReturn struct {
	Order *uc.Order2
	Err   error
}
type GetUserReturn struct {
	User *uc.User2
	Err  error
}

func (tI EvilInterface) GetOrder(id int) (*uc.Order2, error) {
	return tI.GetOrderOutput.Order, tI.GetOrderOutput.Err
}

func (tI EvilInterface) GetUser(id int) (*uc.User2, error) {
	return tI.GetUserOutput.User, tI.GetUserOutput.Err
}
