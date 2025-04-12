package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3197(in io.Reader, out io.Writer) {
	const mod = 100_003
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
	var m, n int
	Fscan(in, &m, &n)
	Fprint(out, (pow(m, n)-m*pow(m-1, n-1)%mod+mod)%mod)
}

//func main() { p3197(os.Stdin, os.Stdout) }
