package tax

import "time"

func CalculateTaxB(amount float64) float64 {
	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}

func CalculateTaxB2(amount float64) float64 {
	if amount == 0 {
		return 0
	}
	time.Sleep(time.Millisecond)
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
