package main

import (
	"fast-gobonacci/fibonacci"
	"fmt"
	"os"
	"strconv"
)

func main() {
	clas := os.Args[1:]
	if i, err := strconv.ParseInt(clas[0], 10, 0); err == nil {
		fib := fibonacci.CalcFib(i)
		fmt.Printf("%d\n", fib)
	}
}
