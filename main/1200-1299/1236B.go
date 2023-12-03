package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1236B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, m int
	Fscan(in, &n, &m)
	Fprint(out, pow(pow(2, m)-1, n))
}

//func main() { cf1236B(os.Stdin, os.Stdout) }
