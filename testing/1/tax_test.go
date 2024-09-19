package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	expected := 2.0
	result := CalculateTax(50.0)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
