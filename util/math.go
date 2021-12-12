package util

func Summation(n float64) float64 {
	var s, i float64

	for i = 1; i <= n; i++ {
		s += i
	}

	return s
}
