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
	tests := []iTester{
		iTester{0, 0},
		iTester{1, 0},
	}

	for k, v := range tests {
		if err := uc.CheckOrder(interfaces.EvilInterface{v.orderOutput, v.userOutput}, 0); err == nil {
			t.Errorf("useCase unable to detect wrong interface return in case #%d", k)
		}
	}
}
