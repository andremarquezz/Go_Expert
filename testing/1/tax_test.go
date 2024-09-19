package tax

import "testing"

// func TestCalculateTax(t *testing.T) {
// 	expected := 2.0
// 	result := CalculateTax(50.0)
// 	if result != expected {
// 		t.Errorf("Expected %f, got %f", expected, result)
// 	}
// }

func TestCalculateTaxBatch(t *testing.T) {
	type test struct {
		amount, expected float64
	}
	table := []test{
		{50.0, 2.0},
		{150.0, 3.0},
		{-50.0, 0.0},
	}
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f, got %f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(150.0)
	}
}
