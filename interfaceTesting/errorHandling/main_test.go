package iTesting_test

import (
	"testing"

	uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling"
	interfaces "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling/mocks"
)

type iTester struct {
	orderOutput int
	userOutput  int
}

func TestCheckOrderUseCase(t *testing.T) {
	GetOrderReturns := []interfaces.GetOrderReturn{
		{&uc.Order{10, 20}, nil},
	}
	GetUserReturns := []interfaces.GetUserReturn{
		{&uc.User{20, "Matth"}, nil},
	}

	for k, v := range GetOrderReturns {
		err := uc.CheckOrder(interfaces.EvilInterface{v, GetUserReturns[0]}, 10)
		check(t, "GetOrder", k, err)
	}
	for k, v := range GetUserReturns {
		err := uc.CheckOrder(interfaces.EvilInterface{GetOrderReturns[0], v}, 10)
		check(t, "GetUser", k, err)
	}
}

func check(t *testing.T, method string, k int, err error) {
	if k == 0 && err != nil {
		t.Errorf("useCase should pass #%d of %s", k, method)
	} else if k != 0 && err == nil {
		t.Errorf("useCase unable to detect wrong interface return in case #%d of %s", k, method)
	}
}
