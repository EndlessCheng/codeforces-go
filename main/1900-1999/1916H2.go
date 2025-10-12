package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://gemini.google.com/app/8a6504a2d3dfb2c8

// https://github.com/EndlessCheng
func cf1916H2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const M = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}
	var n, p, k int
	Fscan(in, &n, &p, &k)

	powPN := pow(p, n)
	x, y, powP := 1, 1, 1
	for range k + 1 {
		Fprint(out, x*pow(y, M-2)%M, " ")
		v := (powPN - powP + M) % M
		x = x * v % M * v % M
		y = y * powP % M * (powP*p%M - 1 + M) % M
		powP = powP * p % M
	}
}

//func main() { cf1916H2(os.Stdin, os.Stdout) }
