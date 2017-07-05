package iTesting_test

import (
	"testing"

	uc "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/initial"
	interfaces "github.com/err0r500/cleanArchitectureGolang/interfaceTesting/initial/mocks"
)

func TestCheckOrderUseCase(t *testing.T) {
	if err := uc.CheckOrder(interfaces.NiceInterface{}, 10); err != nil {
		t.Errorf("an error occured")
	}
}
