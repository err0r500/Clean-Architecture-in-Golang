package iTesting_test

import (
	"errors"
	"testing"

	uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling"
	i "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/errorHandling/mocks"
)

func TestCheckOrderUseCase(t *testing.T) {
	GetOrderReturns := []i.GetOrderReturn{
		{&uc.Order2{10, 20}, nil},
		{&uc.Order2{10, 20}, errors.New("hey")},
		{nil, nil},
		{&uc.Order2{}, nil},
		{&uc.Order2{10, 0}, nil},
	}
	GetUserReturns := []i.GetUserReturn{
		{&uc.User2{20, "Matth"}, nil},
		{&uc.User2{20, "Matth"}, errors.New("text")},
		{nil, nil},
		{&uc.User2{}, nil},
		{&uc.User2{10, "m"}, nil},
	}

	for k, v := range GetOrderReturns {
		err := uc.CheckOrder2(i.EvilInterface{v, GetUserReturns[0]}, 10)
		check(t, "GetOrder", k, err)
	}
	for k, v := range GetUserReturns {
		err := uc.CheckOrder2(i.EvilInterface{GetOrderReturns[0], v}, 10)
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
