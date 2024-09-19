package tax

func CalculateTax(price float64) float64 {
	if price < 0 {
		return 0
	} else if price > 100 {
		return 3.0
	}
	return 2.0
}
