package sellerapi

import "testing"

func TestSimulate(t *testing.T) {

	simulation := Simulate()

	if simulation != "Seller API fake return" {
		t.Errorf("Invalid Simulate() return: %q", simulation)
	}
}
