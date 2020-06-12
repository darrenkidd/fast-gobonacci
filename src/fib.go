package main

import (
	"fast-gobonacci/fibonacci"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func fib(pos int, num string, wg *sync.WaitGroup) {
	defer wg.Done()
	if n, err := strconv.ParseInt(num, 10, 0); err == nil {
		fibOfN := fibonacci.CalcFib(n)
		fmt.Printf("[%d] fib(%d) --> %d\n", pos, n, fibOfN)
	}
}

func main() {
	var wg sync.WaitGroup
	seq := os.Args[1:]
	for i := 0; i < len(seq); i++ {
		wg.Add(1)
		go fib(i+1, seq[i], &wg)
	}
	wg.Wait()
}
