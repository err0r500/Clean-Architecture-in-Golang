package iTesting

import "testing"

type iTester struct {
	orderOutput int
	userOutput  int
}

func TestCheckOrderUseCase(t *testing.T) {
	tests := []iTester{
		iTester{0, 0},
	}

	for k, v := range tests {
		if err := checkOrder(evilInterface{v.orderOutput, v.userOutput}, 0); err == nil {
			t.Errorf("useCase unable to detect wrong interface return in case #%d", k)
		}
	}
}
