package iTesting

import "testing"

func TestCheckOrderUseCase(t *testing.T) {
	if err := checkOrder(niceInterface{}, 10); err != nil {
		t.Errorf("an error occured")
	}
}
