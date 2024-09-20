package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestCalculateTaxWithAssert(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2.0, CalculateTax(50.0))
	assert.Equal(3.0, CalculateTax(150.0))
	assert.Equal(0.0, CalculateTax(-50.0))
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(150.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{50.0, 150.0, -50.0}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount < 0 && result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})
}
