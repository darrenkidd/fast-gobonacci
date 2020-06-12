$ curl "https://www.random.org/sequences/?min=20&max=500&col=1&format=plain&rnd=new" | head -n 5 > sequence.txt 
$ go build fib.go && while read i; do time ./fib $i; done < sequence.txt 
