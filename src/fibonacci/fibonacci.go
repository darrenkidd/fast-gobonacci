package fibonacci

// CalcFib : Calculate the Fibonacci number for n
func CalcFib(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return CalcFib(n-1) + CalcFib(n-2)
}
