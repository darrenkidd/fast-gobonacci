package fibonacci

// import (
// 	"fmt"
// )

var memoisedFibs [100]int64

// CalcFib : Calculate the Fibonacci number for n
func CalcFib(n int64) int64 {
	// fmt.Println("memoised", n, memoisedFibs[n])

	if memoisedFibs[n] > 0 {
		return memoisedFibs[n]
	}

	var thisFib int64

	if n <= 1 {
		thisFib = 1
	} else {
		thisFib = CalcFib(n-1) + CalcFib(n-2)
	}

	memoisedFibs[n] = thisFib
	return thisFib
}
