package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
var mod int

func f(p, c int) int {
	if c == 0 {
		return 1 % mod
	}
	res := (1 + pow(p, (c+1)/2)) * f(p, (c-1)/2)
	if c%2 == 0 {
		res += pow(p, c)
	}
	return res % mod
}

func run(in io.Reader, out io.Writer) {
	var a, x int
	Fscan(in, &a, &x, &mod)
	Fprint(out, f(a, x-1))
}

func main() { run(os.Stdin, os.Stdout) }

func pow(x, n int) (res int) {
	res = 1 % mod
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}
