package main

import (
	"fast-gobonacci/fibonacci"
	"fmt"
	"os"
	"strconv"
)

func fib(pos int, num string) {
	if n, err := strconv.ParseInt(num, 10, 0); err == nil {
		fibOfN := fibonacci.CalcFib(n)
		fmt.Printf("[%d] fib(%d) --> %d\n", pos, n, fibOfN)
	}
}

func main() {
	seq := os.Args[1:]
	for i := 0; i < len(seq); i++ {
		go fib(i+1, seq[i])
	}
	fmt.Scanln()
}
