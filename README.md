## New Way

No fancy stuff; just run everything sequentially:

```bash
$ go build fib.go && time ./fib 42 51 19 27 30
[1] fib(42) --> 433494437
[2] fib(51) --> 32951280099
[3] fib(19) --> 6765
[4] fib(27) --> 317811
[5] fib(30) --> 1346269

real    2m36.423s
user    2m34.083s
sys     0m1.178s
```

## Old Way

```bash
$ curl "https://www.random.org/sequences/?min=20&max=500&col=1&format=plain&rnd=new" | head -n 5 > sequence.txt 
$ go build fib.go && while read i; do time ./fib $i; done < sequence.txt 
...
```
