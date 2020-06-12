## New Way

### `bdafd8d`

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

### `ec2fcf5`

```bash
$ go build fib.go && time ./fib 42 51 19 27 30      
[5] fib(30) --> 1346269
[3] fib(19) --> 6765
[4] fib(27) --> 317811
[1] fib(42) --> 433494437
[2] fib(51) --> 32951280099


real    2m20.442s
user    2m19.272s
sys     0m0.510s
```

### `a194f9a`

```bash
$ go build fib.go && time ./fib 42 51 19 27 30
[1] fib(42) --> 433494437
[2] fib(51) --> 32951280099
[3] fib(19) --> 6765
[4] fib(27) --> 317811
[5] fib(30) --> 1346269

real    0m0.029s
user    0m0.002s
sys     0m0.004s
```

### latest

```bash
$ go build fib.go && time ./fib 42 51 19 27 30
[2] fib(51) --> 32951280099
[3] fib(19) --> 6765
[4] fib(27) --> 317811
[5] fib(30) --> 1346269
[1] fib(42) --> 433494437

real    0m0.371s
user    0m0.004s
sys     0m0.002s
```

## Old Way

```bash
$ curl "https://www.random.org/sequences/?min=20&max=500&col=1&format=plain&rnd=new" | head -n 5 > sequence.txt 
$ go build fib.go && while read i; do time ./fib $i; done < sequence.txt 
...
```
